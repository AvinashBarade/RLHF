package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

// Custom log levels for CI/CD tasks
const (
	logLevelDeploymentSuccess = "DEPLOYMENT_SUCCESS"
	logLevelBuildFailure      = "BUILD_FAILURE"
	logLevelEnvVarChange      = "ENV_VAR_CHANGE"
)

// Function to initialize the logger with a file output
func initLogger(logFile string) error {
	f, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	log.SetOutput(f)
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
	return nil
}

// Function to generate a formatted log message for a CI/CD task
func logTask(level, task, status string, msg ...interface{}) {
	log.Printf("[%s] %s: %s - %s", level, task, status, fmt.Sprint(msg...))
}

func main() {
	// Define the variables for the report
	appVersion := "1.0.2"
	buildTime := time.Now()
	environment := "Staging"

	// Initialize the logger
	logFile := "ci_cd_log.txt"
	if err := initLogger(logFile); err != nil {
		fmt.Printf("Error initializing logger: %v\n", err)
		return
	}

	// Create a formatted string using fmt.Sprintf
	report := fmt.Sprintf("Deployment Report\n----------------\nApp Version: %s\nBuild Time: %s\nEnvironment: %s\n", appVersion, buildTime, environment)
	fmt.Println(report)

	// Example CI/CD tasks and their corresponding log entries
	logTask(logLevelDeploymentSuccess, "Deployment", "Completed", "Application deployed successfully")
	logTask(logLevelBuildFailure, "Build", "Failed", "Build script exited with error: 'Go build failed: cannot find package main'")
	logTask(logLevelEnvVarChange, "Environment", "Updated", "Environment variable 'DB_URL' changed from 'old_url' to 'new_url'")

	fmt.Println("CI/CD tasks completed. Logs saved to", logFile)
}
