package main

import (
	"fmt"
	"os"
	"sync"
)

// Function to simulate a file operation with potential errors
func openFile(filePath string) (*os.File, error) {
	return os.Open(filePath)
}

// Function to simulate processing data from a file with potential errors
func processFile(file *os.File) error {
	// Simulate reading and processing data
	// For demonstration, we will always return an error
	return fmt.Errorf("error processing file")
}

func processFilesConcurrently(filePaths []string) error {
	var wg sync.WaitGroup
	var errors []error

	// Loop over each file path
	for _, filePath := range filePaths {
		wg.Add(1)

		go func(fp string) {
			defer wg.Done()

			file, err := openFile(fp)
			if err != nil {
				errors = append(errors, err)
				return
			}
			defer file.Close() // Ensure the file is closed regardless of the outcome

			err = processFile(file)
			if err != nil {
				errors = append(errors, err)
			}
		}(filePath)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Check if there are any accumulated errors
	if len(errors) > 0 {
		for _, err := range errors {
			fmt.Printf("Error: %v\n", err)
		}
		return fmt.Errorf("some files processing failed")
	}

	return nil
}

func main() {
	filePaths := []string{"example1.txt", "example2.txt", "example3.txt"}

	if err := processFilesConcurrently(filePaths); err != nil {
		fmt.Println("File processing failed:", err)
		return
	}

	fmt.Println("All files processed successfully.")
}
