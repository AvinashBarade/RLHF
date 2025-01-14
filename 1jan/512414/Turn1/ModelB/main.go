package main

import (
	"bytes"
	"fmt"
	"log"
	"sync"
	"time"
)

// Template represents an email template with variations.
type Template struct {
	name        string
	subject     string
	mainBody    *bytes.Buffer // Use bytes.Buffer for efficient appending
	variations  []*bytes.Buffer
	defaultBody int
}

// NewTemplate creates a new Template with a default body.
func NewTemplate(name, subject string, body string) *Template {
	b := bytes.NewBufferString(body)
	return &Template{
		name:        name,
		subject:     subject,
		mainBody:    b,
		variations:  []*bytes.Buffer{b},
		defaultBody: 0,
	}
}

// AddVariation adds a new variation to the template.
func (t *Template) AddVariation(body string) {
	b := bytes.NewBufferString(body)
	t.variations = append(t.variations, b)
}

// SetDefaultBody sets the default body index.
func (t *Template) SetDefaultBody(index int) {
	if index < 0 || index >= len(t.variations) {
		log.Fatalf("Invalid default body index: %d", index)
	}
	t.defaultBody = index
}

// Generate generates the email body with given variables.
func (t *Template) Generate(variables map[string]interface{}) string {
	// Choose a random variation for A/B testing
	index := t.defaultBody
	// For simplicity, we'll use a random number here.
	// In a real A/B testing system, you would use a proper A/B testing library.
	if len(t.variations) > 1 {
		index = int(time.Now().UnixNano() % int64(len(t.variations)))
	}
	b := t.variations[index]

	// Reset the buffer before formatting
	b.Reset()
	// Use fmt.Sprintf for efficient string formatting
	_, err := fmt.Fprintf(b, t.mainBody.String(), variables)
	if err != nil {
		log.Fatalf("Error formatting template: %v", err)
	}
	return b.String()
}

// RenderTemplate renders an email template with given variables.
func RenderTemplate(templateName string, variables map[string]interface{}) string {
	// Look up the template from the template store (omitted for simplicity in this example).
	// Replace with your own template store implementation.
	template, ok := templates[templateName]
	if !ok {
		log.Fatalf("Template not found: %s", templateName)
	}

	// Generate the email body
	body := template.Generate(variables)

	// Construct the email content (omitted for simplicity)
	// For example:
	// emailContent := fmt.Sprintf("Subject: %s\n\n%s", template.subject, body)

	return body
}

// Example usage
var (
	templates map[string]*Template
)

func init() {
	templates = map[string]*Template{}

	// Create the base template
	baseTemplate := NewTemplate("welcome", "Welcome to Our Site!", "Dear %[1]s,\nWelcome to our site!\n\nBest regards,\n[Your Name]")
	templates["welcome"] = baseTemplate

	// Add an A/B test variation
	variation1 := NewTemplate("welcome_variation1", "Welcome to Our Site!", "Dear %[1]s,\nThanks for joining our site!\n\nBest wishes,\n[Your Name]")
	templates["welcome_variation1"] = variation1

	// Add another A/B test variation
	variation2 := NewTemplate("welcome_variation2", "Welcome to Our Site!", "Hey %[1]s!\nWelcome to our community!\n\nWarm regards,\n[Your Name]")
	templates["welcome_variation2"] = variation2