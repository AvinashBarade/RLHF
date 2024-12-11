package main

import (
	"context"
	"io"
	"log"
	"os"
	"time"
)

// processFile reads a file in a goroutine and demonstrates resource cleanup
func processFile(ctx context.Context, filePath string) {
	defer log.Println("File processing goroutine exited")

	file, err := os.Open(filePath)
	if err != nil {
		log.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close() // Ensure file is closed even if an error occurs

	select {
	case <-ctx.Done():
		log.Println("File processing cancelled")
		return
	default:
	}

	log.Printf("Processing file: %s\n", filePath)

	// Simulate file processing by reading it
	_, err = io.Copy(os.Stdout, file)
	if err != nil {
		log.Printf("Error reading file: %v\n", err)
	}

	log.Println("File processing complete")
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel() // Ensure the context is cancelled even if main exits early

	filePath := "example.txt" // Replace with your file path

	go processFile(ctx, filePath)

	log.Println("Main goroutine starting other tasks...")

	time.Sleep(3 * time.Second) // Simulate other work

	log.Println("Cancelling file processing...")
	cancel() // Cancel the context to stop the goroutine

	time.Sleep(2 * time.Second) // Allow goroutine to exit

	log.Println("Main goroutine exiting")
}
