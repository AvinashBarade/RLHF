package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"sync"
	"time"
)

func main() {
	var (
		errorCounter = 0
		mutex        sync.Mutex
	)

	for retry := 0; retry < 3; retry++ { // Number of retry attempts
		file, err := os.Open("example.txt")
		if err != nil {
			log.Printf("Error opening file on attempt %d: %v, will retry.", retry+1, err)
			// Check if the error is "file not found"
			if os.IsNotExist(err) {
				// Only retry if it's "file not found"
				// Wait for 5 seconds before retrying
				time.Sleep(5 * time.Second)
				continue
			} else {
				// Increment error counter if it's not "file not found"
				mutex.Lock()
				errorCounter++
				mutex.Unlock()
			}
		}

		defer file.Close() // Ensure file is closed after successful open

		data, err := io.ReadAll(file)
		if err != nil {
			log.Fatalf("Error reading file: %v", err)
		}
		fmt.Println(string(data))
		break // Exit loop if successful read
	}

	// Check if the loop exited due to successful read or too many errors
	if errorCounter >= 3 {
		log.Printf("Too many errors (%d), giving up.", errorCounter)
	}
}
