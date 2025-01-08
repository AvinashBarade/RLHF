package main

import (
	"bytes"
	"context"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// ProcessLine processes a single line from the CSV file, simulating a long-running operation.
func ProcessLine(ctx context.Context, line string) error {
	// Simulate a long-running task
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

	// Simulate some calculation error
	if _, err := strconv.Atoi(line); err != nil {
		return err
	}

	return nil
}

// ProcessFile reads and processes a CSV file using memory-mapped files and context cancellation.
func ProcessFile(ctx context.Context, filePath string) error {
	// Open the file for memory mapping
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Create a context cancel function
	cancel := ctx.WithCancel()

	// Create a channel to receive errors
	errorsCh := make(chan error, 1)

	// Map the file to memory
	m, err := mmap.Map(file, mmap.Option{})
	if err != nil {
		return err
	}
	defer mmap.Unmap(m)

	// Read the mapped file
	// For simplicity, assume each line is 100 bytes
	fileSize := len(m)
	for i := 0; i < fileSize; i += 100 {
		select {
		case <-cancel.Done():
			fmt.Println("Processing cancelled")
			return cancel.Err()
		default:
			line := m[i : i+100]
			lineStr := string(line)

			// Check for newline characters to determine the actual line size
			newlinePos := bytes.Index(line, []byte{'\n'})
			if newlinePos != -1 {
				lineStr = string(line[:newlinePos])
			}

			// Process the line
			err = ProcessLine(cancel, lineStr)
			if err != nil {
				errorsCh <- err
			}
		}
	}

	// Check for errors in the processing
	select {
	case err := <-errorsCh:
		return err
	default:
		return nil
	}
}

func main() {
	// Set up a new context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	filePath := "large_file.csv"

	err := ProcessFile(ctx, filePath)
	if err != nil {
		fmt.Printf("Error processing file: %v\n", err)
	} else {
		fmt.Println("File processed successfully")
	}
}
