package main

import (
	"fmt"
	"os"
)

// A function to simulate a file operation with potential errors
func openFile(filePath string) (*os.File, error) {
	return os.Open(filePath)
}

// A function to simulate processing data from a file with potential errors
func processFile(file *os.File) error {
	// Simulate reading and processing data
	// For demonstration, we will always return an error
	return fmt.Errorf("error processing file")
}

func main() {
	filePath := "example.txt"

	// Attempt to open the file
	file, err := openFile(filePath)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close() // Ensure the file is closed regardless of the outcome

	// Attempt to process the file
	err = processFile(file)
	if err != nil {
		fmt.Printf("Error processing file: %v\n", err)
		return
	}

	fmt.Println("File processed successfully.")
}
