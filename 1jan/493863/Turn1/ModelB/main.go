//package main

import (
	"fmt"
	"math/rand"
	"time"
)

func slowOperation() (int, error) {
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	if rand.Intn(2) == 0 {
		return 0, fmt.Errorf("random error occurred")
	}
	return 1, nil
}
func main() {
	// Set a flag to indicate whether to enable deferred error handling or not.
	// In a real scenario, you can use command-line flags or configuration settings to control this flag.
	enableDeferredErrorHandling := false
	rand.Seed(time.Now().UnixNano())

	// Perform the slow operation
	result, err := slowOperation()

	// If deferred error handling is enabled, use the 'defer' statement to handle the error.
	if enableDeferredErrorHandling {
		defer func() {
			if err != nil {
				fmt.Println("Error occurred:", err)
			}
		}()
	}

	// Process the result
	if err == nil {
		fmt.Println("Result:", result)
	} else {
		// If deferred error handling is not enabled, handle the error immediately.
		if !enableDeferredErrorHandling {
			fmt.Println("Error occurred:", err)
		}
	}
}
