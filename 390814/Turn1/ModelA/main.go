package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/time/rate"
)

const (
	// Limit - maximum number of requests per second
	Limit        = 2
	BurstLimit   = 5
	RequestCount = 10
)

// Hello represents the function that will execute our requests.
func makeRequest(id int) {
	// Simulate making a request (e.g., to an external API)
	resp, err := http.Get("https://api.github.com")
	if err != nil {
		fmt.Printf("Request %d failed: %v\n", id, err)
		return
	}
	defer resp.Body.Close()
	fmt.Printf("Request %d succeeded: %s\n", id, resp.Status)
}

func main() {
	// Create a rate limiter that allows 2 requests per second with a burst size of 5.
	limiter := rate.NewLimiter(rate.Limit(Limit), BurstLimit)

	// Use a channel to synchronize concurrent execution.
	sem := make(chan struct{}, BurstLimit)

	for i := 1; i <= RequestCount; i++ {
		sem <- struct{}{} // Acquire a slot in the semaphore.
		go func(id int) {
			defer func() { <-sem }() // Release the slot when function completes.

			// Wait until we can proceed according to the limit.
			if err := limiter.Wait(context.Background()); err != nil {
				fmt.Printf("Request %d was not allowed: %v\n", id, err)
				return
			}

			makeRequest(id)
		}(i)
	}

	// Wait for a while to allow all goroutines to finish for demonstration purpose
	time.Sleep(5 * time.Second)
}
