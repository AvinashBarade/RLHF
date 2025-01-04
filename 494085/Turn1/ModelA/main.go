package main

import (
	"fmt"
	"os"
)

func openFile(filename string) (*os.File, error) {
	fmt.Printf("Opening file: %s\n", filename)
	return os.Open(filename)
}

func readFile(file *os.File) error {
	fmt.Printf("Reading file...\n")
	// Simulate an error by causing a panic
	panic("Error reading file!")
}

func closeFile(file *os.File) {
	fmt.Printf("Closing file...\n")
	if err := file.Close(); err != nil {
		fmt.Printf("Error closing file: %v\n", err)
	}
}

func main() {
	// Defer is used to ensure the file is closed, even if a panic occurs
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from panic: %v\n", r)
		}
	}()

	file, err := openFile("example.txt")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer closeFile(file)

	err = readFile(file)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
	}
}
