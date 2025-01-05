package main

import (
	"fmt"
	"os"
	"time"
)

// LogEntry represents a single log entry
type LogEntry struct {
	Timestamp time.Time
	TaskType  string
	Status    string
	Details   string
}

// Log specifies the file name and appends log entries to it
type Log struct {
	file *os.File
}

// NewLog creates a new Log object that appends log entries to a given file
func NewLog(filename string) (*Log, error) {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return nil, fmt.Errorf("failed to open log file: %v", err)
	}
	return &Log{file: file}, nil
}

// Write writes a LogEntry to the log file
func (l *Log) Write(entry LogEntry) error {
	// Format the log entry
	formattedEntry := fmt.Sprintf("[%s] %s - %s: %s\n", entry.Timestamp.Format(time.RFC3339), entry.TaskType, entry.Status, entry.Details)
	_, err := l.file.WriteString(formattedEntry)
	return err
}

// Close closes the underlying file of the Log
func (l *Log) Close() error {
	return l.file.Close()
}

func main() {
	// Create a Log instance
	logFileName := "ci_cd_logs.txt"
	log, err := NewLog(logFileName)
	if err != nil {
		fmt.Printf("Error creating log: %v\n", err)
		return
	}
	defer log.Close() // Ensure log is closed at the end

	// Example log entries
	var logEntries []LogEntry
	logEntries = append(logEntries, LogEntry{
		Timestamp: time.Now(),
		TaskType:  "Deployment",
		Status:    "Success",
		Details:   "Deployed version 1.0.1 to Production",
	})
	logEntries = append(logEntries, LogEntry{
		Timestamp: time.Now().Add(-time.Hour),
		TaskType:  "Build",
		Status:    "Failed",
		Details:   "Test coverage failed: 85% required, got 80%",
	})
	logEntries = append(logEntries, LogEntry{
		Timestamp: time.Now().Add(-time.Minute),
		TaskType:  "Environment Change",
		Status:    "Succeeded",
		Details:   "Updated ENV_KEY=new_value in Production environment",
	})

	// Write log entries
	for _, entry := range logEntries {
		if err := log.Write(entry); err != nil {
			fmt.Printf("Error writing log entry: %v\n", err)
		}
	}

	fmt.Println("CI/CD task logs saved successfully!")
}
