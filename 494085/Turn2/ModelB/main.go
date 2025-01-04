package main

import (
	"fmt"
	"log"
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

func closeFile(file *os.File, logger *log.Logger) {
	fmt.Printf("Closing file...\n")
	if err := file.Close(); err != nil {
		logger.Printf("Error closing file: %v\n", err)
	} else {
		logger.Println("File closed successfully.")
	}
}

func main() {
	// Create a logger for logging to defer_log.txt
	logFile, err := os.Create("defer_log.txt")
	if err != nil {
		fmt.Printf("Error creating log file: %v\n", err)
		return
	}
	defer logFile.Close()
	logger := log.New(logFile, "", log.LstdFlags)

	// Defer is used to ensure the file is closed, even if a panic occurs
	defer func() {
		if r := recover(); r != nil {
			logger.Printf("Recovered from panic: %v\n", r)
		}
	}()

	file, err := openFile("example.txt")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer closeFile(file, logger)

	err = readFile(file)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
	}
}
