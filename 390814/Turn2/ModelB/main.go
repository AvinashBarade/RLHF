package main

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"golang.org/x/time/rate"
)

const (
	// Limit - maximum number of requests per second.
	Limit = 2
	// BurstLimit - burst size for rate limiter.
	BurstLimit = 5
	// RequestCount - total number of requests to make.
	RequestCount = 10
	// MaxRetries - maximum number of retries for each request.
	MaxRetries = 3
	// BackoffBase - base duration for exponential backoff.
	BackoffBase = 100 * time.Millisecond
)

// makeRequest sends an HTTP request and handles retries with exponential backoff.
func makeRequest(id int, limiter *rate.Limiter) {
	// Acquire a slot in the rate limiter.
	if err := limiter.Wait(context.Background()); err != nil {
		fmt.Printf("Request %d was not allowed: %v\n", id, err)
		return
	}

	for retry := 0; retry <= MaxRetries; retry++ {
		// Simulate making an HTTP request.
		resp, err := http.Get("https://api.github.com")
		if err == nil {
			defer resp.Body.Close()
			fmt.Printf("Request %d succeeded: %s\n", id, resp.Status)
			return
		}

		// Handle request failure with exponential backoff.
		fmt.Printf("Request %d failed: %v. Retrying...\n", id, err)

		// Calculate exponential backoff delay.
		backoff := time.Duration(rand.Intn(1000)) * BackoffBase * time.Duration(retry)
		time.Sleep(backoff)
	}

	// If all retries fail, log the final error.
	fmt.Printf("Request %d failed after %d retries.\n", id, MaxRetries)
}

func main() {
	// Create a rate limiter with the specified limits.
	limiter := rate.NewLimiter(rate.Limit(Limit), BurstLimit)

	for i := 1; i <= RequestCount; i++ {
		// Launch a goroutine to make the request with retry logic.
		go makeRequest(i, limiter)
	}

	// Wait for a while to allow all goroutines to finish for demonstration purpose.
	time.Sleep(5 * time.Second)
}
