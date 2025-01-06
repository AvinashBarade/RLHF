package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"sync"
	"time"
)

const maxRetryAttempts = 5                       // Max retry attempts
const initialBackoffTime = 50 * time.Millisecond // Initial backoff time

func main() {
	var (
		errorCounter = 0
		mutex        sync.Mutex
	)

	file, err := os.Open("example.txt")
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	data, err := readFileWithRetry(file, &errorCounter)
	if err != nil {
		log.Printf("Giving up after %d retries: %v", maxRetryAttempts, err)
	} else {
		fmt.Println(string(data))
	}
}

func readFileWithRetry(file *os.File, errorCounter *int) ([]byte, error) {
	for attempt := 1; attempt <= maxRetryAttempts; attempt++ {
		data, err := io.ReadAll(file)
		if err != nil {
			log.Printf("Error reading file on attempt %d: %v", attempt, err)
			file.Seek(0, 0) // Rewind the file to read from the beginning on the next attempt
			mutex.Lock()
			*errorCounter++
			mutex.Unlock()
			if *errorCounter >= 3 {
				log.Printf("Too many errors (%d), stopping retries.", *errorCounter)
				return nil, err
			}
			// Wait for a random amount of time before retrying with backoff
			backoff := backoffDuration(attempt)
			log.Printf("Retrying after %v", backoff)
			time.Sleep(backoff)
		} else {
			return data, nil // Success!
		}
	}
	return nil, fmt.Errorf("exceeded max retry attempts")
}

func backoffDuration(attempt int) time.Duration {
	// Calculate exponential backoff with a random jitter to avoid hammering the file system
	base := float64(initialBackoffTime)
	backoff := time.Duration(base*(2.0**float64(attempt - 1))) + time.Duration(rand.Int63n(int64(base)))
	return backoff
}
