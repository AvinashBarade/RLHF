package main

import (
	"fmt"
	"os"
)

func main() {
	// Example: using fmt.Fprintf to write to a file
	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		return
	}
	defer file.Close()

	if _, err := fmt.Fprintf(file, "Hello, %s!\n", "Gopher"); err != nil {
		// Handle I/O error writing to file
		fmt.Printf("Error writing to file: %v\n", err)
		return
	}

	// Example: using fmt.Errorf for returning a formatted error
	userID := 42
	if userID < 0 {
		err := fmt.Errorf("userID cannot be negative: %d", userID)
		// Handle this error
		fmt.Println("Error:", err)
	}

	// Meanwhile, fmt.Sprintf never returns an error
	msg := fmt.Sprintf("User ID: %d", userID)
	fmt.Println(msg) // "User ID: 42"
}
