package main

import (
	"fmt"
	"sync"
	"time"
)

// function to simulate some work and measure execution time
func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()         // Mark this Goroutine as done when the function returns
	startTime := time.Now() // Record the start time

	// Simulate work (e.g., sleep for a random duration)
	time.Sleep(time.Duration(id) * 100 * time.Millisecond)

	// Calculate execution time
	elapsedTime := time.Since(startTime)
	defer reportExecutionTime(id, elapsedTime) // Ensure we report execution time at the end

	// Additional processing can be done here
}

// function to report execution time
func reportExecutionTime(id int, duration time.Duration) {
	fmt.Printf("Goroutine %d finished in %s\n", id, duration)
}

func main() {
	var wg sync.WaitGroup
	numWorkers := 5

	// Launch multiple Goroutines
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)         // Increment the WaitGroup counter
		go worker(i, &wg) // Start the worker Goroutine
	}

	wg.Wait() // Wait for all Goroutines to finish
	fmt.Println("All Goroutines have completed.")
}
