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

func closeFile(file *os.File, logFile *os.File) {
	defer func() {
		log.Printf("Closing file...\n")
		if err := file.Close(); err != nil {
			logFile.WriteString(fmt.Sprintf("Error closing file: %v\n", err))
			log.Printf("Error closing file: %v\n", err)
		} else {
			logFile.WriteString("File closed successfully\n")
			log.Printf("File closed successfully\n")
		}
	}()
}

func main() {
	// Set up logging to defer_log.txt
	logFile, err := os.Create("defer_log.txt")
	if err != nil {
		fmt.Printf("Error creating log file: %v\n", err)
		return
	}
	defer logFile.Close()

	// Defer is used to ensure the file is closed, even if a panic occurs
	defer func() {
		if r := recover(); r != nil {
			logFile.WriteString(fmt.Sprintf("Recovered from panic: %v\n", r))
			fmt.Printf("Recovered from panic: %v\n", r)
		}
	}()

	file, err := openFile("example.txt")
	if err != nil {
		logFile.WriteString(fmt.Sprintf("Error opening file: %v\n", err))
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer closeFile(file, logFile)

	err = readFile(file)
	if err != nil {
		logFile.WriteString(fmt.Sprintf("Error reading file: %v\n", err))
		fmt.Printf("Error reading file: %v\n", err)
	}
}
