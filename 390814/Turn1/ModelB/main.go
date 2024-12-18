package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"

	"github.com/hashicorp/limiter"
)

// Constants for rate limiting
const (
	limit    = 5  // Requests per second
	burst    = 10 // Allow bursting up to 10 requests
	duration = 1 * time.Second
)

// Create a rate limiter with the specified rate and burst
var rl *limiter.Limiter

func init() {
	var err error
	rl, err = limiter.New(limiter.Every(duration), limit, limiter.WithBurst(burst))
	if err != nil {
		log.Fatalf("Failed to create limiter: %v", err)
	}
}

func simulateWork() {
	fmt.Println("Starting work...")
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	fmt.Println("Work completed.")
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			// Check if the request is allowed
			if err := rl.Take(context.Background()); err != nil {
				fmt.Println("Request rate limit exceeded:", err)
				return
			}

			simulateWork()
		}()
	}

	wg.Wait()
	fmt.Println("All work completed.")
}
