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

	file, err := os.Open("example.txt")
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer func() {
		mutex.Lock()
		defer mutex.Unlock()
		errorCounter++
		if errorCounter < 3 {
			file.Close()
		} else {
			// Adaptive behavior: Stop retrying after 3 errors
			log.Printf("Too many errors (%d), stopping retries.", errorCounter)
		}
	}()

	// Retry logic for reading the file
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		data, err := io.ReadAll(file)
		if err != nil {
			log.Printf("Error reading file: %v, retrying attempt %d.", err, i+1)
			// Increment error counter and check for retry limit
			mutex.Lock()
			errorCounter++
			mutex.Unlock()
			if errorCounter >= 3 {
				log.Printf("Too many errors (%d), giving up.", errorCounter)
				return
			}
			// Sleep before retrying
			time.Sleep(time.Second)
		} else {
			fmt.Println(string(data))
			break // Exit loop if successful read
		}
	}
}
