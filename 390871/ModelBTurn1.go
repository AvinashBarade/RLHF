package main

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

// maxRetries defines the maximum number of retries
const maxRetries = 3

// retryDelayBase is the base delay for exponential backoff
const retryDelayBase = time.Second

// performRequest performs an HTTP request with retry logic
func performRequest(ctx context.Context, url string) error {
	for retryCount := 0; retryCount <= maxRetries; retryCount++ {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		resp, err := http.Get(url)
		if err == nil {
			defer resp.Body.Close()
			if resp.StatusCode == http.StatusOK {
				fmt.Println("Request successful")
				return nil
			}
			fmt.Printf("Request failed with status %d\n", resp.StatusCode)
		} else {
			fmt.Printf("Request failed with error: %v\n", err)
		}

		// Calculate delay with jitter
		delay := retryDelayBase * time.Duration(1<<uint(retryCount))
		jitter := time.Duration(rand.Intn(int(delay.Nanoseconds() / 1e6)))
		delay += jitter

		fmt.Printf("Retrying in %v...\n", delay)
		select {
		case <-time.After(delay):
		case <-ctx.Done():
			return ctx.Err()
		}
	}

	return fmt.Errorf("max retries reached")
}

func main() {
	rand.Seed(time.Now().UnixNano())
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	go func() {
		if err := performRequest(ctx, "https://example.com/nonexistent"); err != nil {
			fmt.Println("Failed to perform request:", err)
		}
	}()

	// Simulate work
	time.Sleep(5 * time.Second)
}
