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
	maxGoroutines = 10 // Limitation for the first scheduling pattern
	url           = "https://jsonplaceholder.typicode.com/todos/1"
)

type ResponseMetrics struct {
	SuccessCount     int
	FailureCount     int
	TotalErrorTime   time.Duration
	AverageErrorTime time.Duration
}

// Function to perform an HTTP GET request and return metrics
func fetch(url string, wg *sync.WaitGroup, results chan<- string, metrics *ResponseMetrics) {
	defer wg.Done()

	start := time.Now()
	resp, err := http.Get(url)

	if err != nil {
		metrics.FailureCount++
		metrics.TotalErrorTime += time.Since(start)
		results <- fmt.Sprintf("Error: %v", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		metrics.FailureCount++
		metrics.TotalErrorTime += time.Since(start)
		results <- fmt.Sprintf("Error reading response: %v", err)
		return
	}

	metrics.SuccessCount++
	results <- fmt.Sprintf("Response size: %d bytes", len(body))
}

// Function to perform concurrent requests with a limit on Goroutines
func limitedConcurrency() {
	var wg sync.WaitGroup
	results := make(chan string, totalRequests)
	metrics := &ResponseMetrics{}

	sem := make(chan struct{}, maxGoroutines) // Semaphore to limit concurrency

	start := time.Now()
	for i := 0; i < totalRequests; i++ {
		sem <- struct{}{} // Acquire a token
		wg.Add(1)

		go func() {
			defer func() { <-sem }() // Release the token
			fetch(url, &wg, results, metrics)
		}()
	}

	wg.Wait()
	close(results)

	// Calculate average error time
	if metrics.FailureCount > 0 {
		metrics.AverageErrorTime = metrics.TotalErrorTime / time.Duration(metrics.FailureCount)
	}

	elapsed := time.Since(start)
	printResults(metrics, elapsed)
}

// Function to perform concurrent requests without limiting Goroutines
func unlimitedConcurrency() {
	var wg sync.WaitGroup
	results := make(chan string, totalRequests)
	metrics := &ResponseMetrics{}

	start := time.Now()
	for i := 0; i < totalRequests; i++ {
		wg.Add(1)
		go fetch(url, &wg, results, metrics)
	}

	wg.Wait()
	close(results)

	// Calculate average error time
	if metrics.FailureCount > 0 {
		metrics.AverageErrorTime = metrics.TotalErrorTime / time.Duration(metrics.FailureCount)
	}

	elapsed := time.Since(start)
	printResults(metrics, elapsed)
}

// Function to print the results and metrics
func printResults(metrics *ResponseMetrics, elapsed time.Duration) {
	fmt.Printf("Total Time: %s\n", elapsed)
	fmt.Printf("Successful Requests: %d\n", metrics.SuccessCount)
	fmt.Printf("Failed Requests: %d\n", metrics.FailureCount)
	if metrics.FailureCount > 0 {
		fmt.Printf("Average Time for Failed Requests: %s\n", metrics.AverageErrorTime)
	} else {
		fmt.Println("No Failed Requests.")
	}
}

func main() {
	fmt.Println("=== Limited Concurrency ===")
	limitedConcurrency()

	fmt.Println("\n=== Unlimited Concurrency ===")
	unlimitedConcurrency()
}
