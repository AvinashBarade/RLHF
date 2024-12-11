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

// Function to perform an HTTP GET request
func fetch(url string, wg *sync.WaitGroup, results chan<- string) {
	defer wg.Done()
	resp, err := http.Get(url)
	if err != nil {
		results <- fmt.Sprintf("Error: %v", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		results <- fmt.Sprintf("Error reading response: %v", err)
		return
	}
	results <- fmt.Sprintf("Response size: %d bytes", len(body))
}

// Function to perform concurrent requests with a limit on Goroutines
func limitedConcurrency() {
	var wg sync.WaitGroup
	results := make(chan string, totalRequests)

	sem := make(chan struct{}, maxGoroutines) // Semaphore to limit concurrency

	start := time.Now()
	for i := 0; i < totalRequests; i++ {
		sem <- struct{}{} // Acquire a token
		wg.Add(1)

		go func() {
			defer func() { <-sem }() // Release the token
			fetch(url, &wg, results)
		}()
	}

	wg.Wait()
	close(results)
	elapsed := time.Since(start)

	fmt.Printf("Limited Concurrency took: %s\n", elapsed)
	for res := range results {
		fmt.Println(res)
	}
}

// Function to perform concurrent requests without limiting Goroutines
func unlimitedConcurrency() {
	var wg sync.WaitGroup
	results := make(chan string, totalRequests)

	start := time.Now()
	for i := 0; i < totalRequests; i++ {
		wg.Add(1)
		go fetch(url, &wg, results)
	}

	wg.Wait()
	close(results)
	elapsed := time.Since(start)

	fmt.Printf("Unlimited Concurrency took: %s\n", elapsed)
	for res := range results {
		fmt.Println(res)
	}
}

func main() {
	fmt.Println("=== Limited Concurrency ===")
	limitedConcurrency()

	fmt.Println("\n=== Unlimited Concurrency ===")
	unlimitedConcurrency()
}
