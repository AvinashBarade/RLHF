package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

const (
	totalRequests = 100
	maxGoroutines = 10 // Limit for the first scheduling pattern
	url           = "https://jsonplaceholder.typicode.com/todos/1"
)

type Result struct {
	responseSize int
	err          error
	duration     time.Duration
}

// Function to perform an HTTP GET request
func fetch(url string, wg *sync.WaitGroup, results chan<- Result) {
	defer wg.Done()
	start := time.Now()

	resp, err := http.Get(url)
	duration := time.Since(start)

	if err != nil {
		results <- Result{err: err, duration: duration}
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		results <- Result{err: err, duration: duration}
		return
	}
	results <- Result{responseSize: len(body), duration: duration}
}

// Function to perform concurrent requests with a limit on Goroutines
func limitedConcurrency() {
	var wg sync.WaitGroup
	results := make(chan Result, totalRequests)

	sem := make(chan struct{}, maxGoroutines) // Semaphore to limit concurrency
	var failedRequests int
	var totalFailedDuration time.Duration

	start := time.Now()
	for i := 0; i < totalRequests; i++ {
		sem <- struct{}{} // Acquire a token
		wg.Add(1)

		go func() {
			defer func() { <-sem }() // Release the token
			result := fetch(url, &wg, results)

			if result.err != nil {
				failedRequests++
				totalFailedDuration += result.duration
			}
		}()
	}

	wg.Wait()
	close(results)
	elapsed := time.Since(start)

	avgFailedDuration := time.Duration(0)
	if failedRequests > 0 {
		avgFailedDuration = totalFailedDuration / time.Duration(failedRequests)
	}

	fmt.Printf("Limited Concurrency took: %s\n", elapsed)
	fmt.Printf("Failed Requests: %d\n", failedRequests)
	fmt.Printf("Average Failed Request Duration: %s\n", avgFailedDuration)
}

// Function to perform concurrent requests without limiting Goroutines
func unlimitedConcurrency() {
	var wg sync.WaitGroup
	results := make(chan Result, totalRequests)
	var failedRequests int
	var totalFailedDuration time.Duration

	start := time.Now()
	for i := 0; i < totalRequests; i++ {
		wg.Add(1)
		go fetch(url, &wg, results)
	}

	wg.Wait()
	close(results)
	elapsed := time.Since(start)

	avgFailedDuration := time.Duration(0)
	if failedRequests > 0 {
		avgFailedDuration = totalFailedDuration / time.Duration(failedRequests)
	}

	fmt.Printf("Unlimited Concurrency took: %s\n", elapsed)
	fmt.Printf("Failed Requests: %d\n", failedRequests)
	fmt.Printf("Average Failed Request Duration: %s\n", avgFailedDuration)
}

func main() {
	fmt.Println("=== Limited Concurrency ===")
	limitedConcurrency()

	fmt.Println("\n=== Unlimited Concurrency ===")
	unlimitedConcurrency()
}
