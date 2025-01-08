package main

import (
	"fmt"
	"log"
	"os"
)

// FilterPositiveNumbers filters out negative integers from a slice of mixed types while skipping non-integer values.
func FilterPositiveNumbers(numbers interface{}) ([]int, error) {
	// Check for nil slice
	if numbers == nil {
		log.Println("ERROR: Input slice is nil")
		return nil, fmt.Errorf("input slice is nil")
	}

	// Check if the input is a slice
	sliceValue, ok := numbers.([]interface{})
	if !ok {
		log.Println("ERROR: Input is not a slice")
		return nil, fmt.Errorf("input is not a slice")
	}

	// Filter positive integers
	filtered := []int{}
	for index, element := range sliceValue {
		switch element := element.(type) {
		case int:
			if element >= 0 {
				filtered = append(filtered, element)
			} else {
				log.Printf("Skipped negative integer: %d at index %d\n", element, index)
			}
		case float64:
			if element > 0 {
				log.Printf("Skipped float: %.2f at index %d\n", element, index)
			}
		default:
			log.Printf("Skipped non-integer value: %v of type %T at index %d\n", element, element, index)
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

	// Example input slice with mixed types
	input := []interface{}{10, -5, "hello", 0, 23.5, -7, 8, "world", 3.14, -4}

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
