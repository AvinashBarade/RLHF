package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

var (
	templates map[string]string
)

func init() {
	// Initialize template cache and generate a random seed for A/B testing
	templates = make(map[string]string)
	rand.Seed(time.Now().UnixNano())
}

func loadTemplate(name string, path string) (string, error) {
	// Check if template is already loaded
	if template, ok := templates[name]; ok {
		return template, nil
	}

	// Read the template from a file
	filePath := filepath.Join(path, name+".txt")
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to read template file: %s", err)
	}

	// Store the loaded template
	templates[name] = string(data)
	return string(data), nil
}

func generateEmail(templateName string, vars map[string]interface{}) (string, error) {
	// Determine which template variation to use for A/B testing
	var variant string
	switch templateName {
	case "welcome":
		variant = "a" // Define default variation
		if rand.Float64() < 0.5 {
			variant = "b"
		}
		templateName = fmt.Sprintf("%s_%s", templateName, variant)
	}

	// Load the template variation
	template, err := loadTemplate(templateName, "templates")
	if err != nil {
		return "", err
	}

	// Format the template with variables
	var formatted string
	var startTime time.Time = time.Now()
	defer func() {
		duration := time.Now().Sub(startTime)
		log.Printf("Formatted email %q in %v\n", templateName, duration)
	}()

	formatted, err = fmt.Sprintf(template, vars...)
	if err != nil {
		return "", fmt.Errorf("failed to format template %q with variables: %v", templateName, err)
	}

	return formatted, nil
}

func main() {
	vars := map[string]interface{}{
		"name":     "Alice",
		"product":  "Wonderful Widget",
		"discount": "15%",
	}

	email, err := generateEmail("welcome", vars)
	if err != nil {
		log.Fatalf("Failed to generate email: %v\n", err)
	}

	log.Printf("Generated Email:\n%s\n", email)
}
