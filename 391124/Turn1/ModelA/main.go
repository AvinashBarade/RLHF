package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

// Mocking a third-party service call
func callThirdPartyService(ctx context.Context, url string, wg *sync.WaitGroup, results chan<- string) {
	defer wg.Done()

	// Create a new HTTP client with a timeout
	client := http.Client{
		Timeout: 5 * time.Second, // Set the timeout for each request
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		results <- fmt.Sprintf("Failed to create request: %v", err)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		results <- fmt.Sprintf("Failed to call service: %v", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		results <- fmt.Sprintf("Request failed with status: %s", resp.Status)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		results <- fmt.Sprintf("Failed to read response body: %v", err)
		return
	}

	results <- string(body) // Pass the result back to the channel
}

func main() {
	// Example URLs for demonstration
	urls := []string{
		"https://jsonplaceholder.typicode.com/posts/1",
		"https://jsonplaceholder.typicode.com/posts/2",
		"https://jsonplaceholder.typicode.com/posts/3",
	}

	var wg sync.WaitGroup
	results := make(chan string, len(urls)) // Channel to gather results

	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	for _, url := range urls {
		wg.Add(1)                                        // increase wait group counter
		go callThirdPartyService(ctx, url, &wg, results) // Launch goroutine
	}

	// Wait for all Goroutines to finish
	go func() {
		wg.Wait()
		close(results) // close results channel after all are done
	}()

	// Collect results
	for result := range results {
		fmt.Println(result)
	}
}
