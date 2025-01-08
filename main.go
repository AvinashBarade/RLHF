package main

import (
	"fmt"
	"log"
	"os"
)

// FilterPositiveNumbers filters out negative numbers from a slice of integers while maintaining order.
func FilterPositiveNumbers(numbers []int) ([]int, error) {
	// Check for nil slice
	if numbers == nil {
		log.Println("ERROR: Input slice is nil")
		return nil, fmt.Errorf("input slice is nil")
	}

	// Filter positive numbers
	filtered := []int{}
	for _, num := range numbers {
		if num >= 0 {
			filtered = append(filtered, num)
		}
	}

	return filtered, nil
}

func setupLogger() {
	// Configure the logger with custom settings
	logFile, err := os.OpenFile("application.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Failed to set up log file: %v\n", err)
		os.Exit(1)
	}
	log.SetOutput(logFile)
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.SetPrefix("[FilterLogger] ")
}

func main() {
	// Set up the logger
	setupLogger()

	// Example input slice
	input := []int{10, -5, 0, 23, -7, 8}

	log.Println("Starting filtering process")

	// Process the slice
	filtered, err := FilterPositiveNumbers(input)
	if err != nil {
		log.Printf("ERROR: %v\n", err)
		return
	}

	log.Printf("Filtering complete. Input: %v, Filtered: %v\n", input, filtered)

	// Output the result
	fmt.Println("Filtered slice:", filtered)
}
