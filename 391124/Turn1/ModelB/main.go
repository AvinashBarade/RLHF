package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

// mockThirdPartyCall simulates an HTTP call to a third-party service
func mockThirdPartyCall(ctx context.Context, url string, wg *sync.WaitGroup) {
	defer wg.Done()
	select {
	case <-ctx.Done():
		log.Printf("Context canceled for URL: %s\n", url)
		return
	default:
	}

	// Simulate HTTP request
	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)

	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error fetching %s: %v\n", url, err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Failed to get %s: %d\n", url, resp.StatusCode)
		return
	}

	log.Printf("Successfully fetched %s\n", url)
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	var wg sync.WaitGroup
	urls := []string{
		"https://jsonplaceholder.typicode.com/posts/1",
		"https://jsonplaceholder.typicode.com/posts/2",
		"https://jsonplaceholder.typicode.com/posts/5",
		"https://jsonplaceholder.typicode.com/posts/5",
		"https://jsonplaceholder.typicode.com/posts/5",
	}

	// Launch goroutines for each URL
	for _, url := range urls {
		wg.Add(1)
		go mockThirdPartyCall(ctx, url, &wg)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	fmt.Println("All requests completed.")
}
