package main

import (
	"fmt"
	"os"
)

func openFile(filename string) *os.File {
	file, err := os.Open(filename)
	if err != nil {
		panic(fmt.Sprintf("failed to open file: %v", err))
	}
	return file
}

func writeToFile(file *os.File) {
	defer file.Close() // defer statement to close the file
	// do some file writing operations
	fmt.Fprintln(file, "Hello, world!")
	// Simulate a panic
	panic("Oops! An error occurred.")
}

func main() {
	defer fmt.Println("Program exiting.")

	filename := "example.txt"
	fmt.Println("Attempting to write to file", filename)

	file := openFile(filename)

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
			fmt.Println("Performing cleanup operations...")
			// Add additional cleanup tasks here
		}
	}()

	writeToFile(file)
	// If the program exits normally, the "defer" statements will be executed in last-in-first-out (LIFO) order.
	fmt.Println("Data written to file successfully.")
}
