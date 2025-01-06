package main

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

func main() {
	var (
		errorCounter = 0
		mutex        = sync.Mutex{}
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

	// Simulate a low-latency system with frequent errors
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		data, err := os.ReadAll(file)
		if err != nil {
			log.Printf("Error reading file: %v", err)
			// Retry logic can be added here
		} else {
			fmt.Println(string(data))
		}
	}
}
