package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"sync"
	"time"
)

const maxRetries = 3

func main() {
	var (
		errorCounter = 0
		mutex        sync.Mutex
	)
	// Retry loop for opening the file
	for i := 0; i < maxRetries; i++ {
		file, err := os.Open("example.txt")
		if err == nil {
			defer func() {
				mutex.Lock()
				defer mutex.Unlock()
				errorCounter++
				if errorCounter < maxRetries {
					file.Close()
				}
			}()

			// Retry logic for reading the file
			for j := 0; j < 5; j++ {
				time.Sleep(100 * time.Millisecond)
				data, err := io.ReadAll(file)
				if err != nil {
					log.Printf("Error reading file: %v, retrying attempt %d.", err, j+1)
					// Increment error counter and check for retry limit
					mutex.Lock()
					errorCounter++
					mutex.Unlock()
					if errorCounter >= maxRetries {
						log.Printf("Too many errors (%d), giving up.", errorCounter)
						return
					}
					// Sleep before retrying
					time.Sleep(time.Second)
				} else {
					fmt.Println(string(data))
					return // Exit loop if successful read
				}
			}
		}
		// File not found error case
		if os.IsNotExist(err) {
			log.Printf("File not found. Retrying in 5 seconds attempt %d.", i+1)
			time.Sleep(5 * time.Second)
			continue
		}

		// Other error case
		log.Printf("Unable to open file, error: %v", err)
		return
	}

	log.Printf("Maximum retries reached. Giving up.")
}
