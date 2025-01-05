package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	fileName := "example.txt"

	if err := handleFileError(fileName); err != nil {
		fmt.Println(err)
	}
}

func handleFileError(fileName string) error {
	if _, err := os.Stat(fileName); err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("file not found: %w", err)
		}

		// Check if the error is related to permissions
		if pathErr, ok := err.(*os.PathError); ok {
			if pathErr.Err == syscall.EACCES {
				return fmt.Errorf("permission denied to access file: %w", err)
			}
		}

		// Return the generic error if it's not specific
		return fmt.Errorf("an error occurred while accessing the file: %w", err)
	}

	return nil // File exists and has appropriate permissions
}
