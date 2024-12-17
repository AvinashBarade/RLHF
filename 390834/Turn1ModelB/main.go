package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"time"
)

// ReadFileWithContext reads a file within the given context.
// It will return an error if the context is cancelled or the deadline is exceeded.
func ReadFileWithContext(ctx context.Context, filename string) ([]byte, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err() // Return the context error if it's cancelled or deadline is exceeded
	default:
	}

	// Proceed with reading the file
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// WriteFileWithContext writes data to a file within the given context.
// It will return an error if the context is cancelled or the deadline is exceeded.
func WriteFileWithContext(ctx context.Context, filename string, data []byte) error {
	select {
	case <-ctx.Done():
		return ctx.Err() // Return the context error if it's cancelled or deadline is exceeded
	default:
	}

	// Proceed with writing the file
	err := ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	// Create a new context with a deadline 2 seconds from now
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel() // Clean up the context when done

	// Simulate a file operation that might take some time
	time.Sleep(1 * time.Second)

	// Read a file with context
	data, err := ReadFileWithContext(ctx, "example.txt")
	if err != nil {
		log.Printf("Error reading file: %v\n", err)
		return
	}

	fmt.Println("File content:")
	fmt.Println(string(data))

	// Write the data to another file with context
	err = WriteFileWithContext(ctx, "output.txt", data)
	if err != nil {
		log.Printf("Error writing file: %v\n", err)
		return
	}

	fmt.Println("File written successfully.")
}
