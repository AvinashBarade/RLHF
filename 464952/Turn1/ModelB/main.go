package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("example.txt") // Open a file for reading
	if err != nil {                     // Check if there was an error opening the file
		fmt.Println("Error opening file:", err)
		return
	}

	defer file.Close() // Ensure the file is closed after the main function returns

	// Read the file content
	content := make([]byte, 1024)
	n, err := file.Read(content)
	if err != nil { // Check if there was an error reading the file
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Printf("Read %d bytes:\n%s", n, string(content[:n]))
}
