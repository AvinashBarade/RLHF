package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

// Initialize logger
var logger = logrus.New()

// Struct for API call results
type APICallResult struct {
	URL      string
	Status   string
	Response string
	Error    error
}

func callThirdPartyService(ctx context.Context, url string, wg *sync.WaitGroup, results chan<- APICallResult) {
	defer wg.Done()

	client := http.Client{
		Timeout: 5 * time.Second,
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		results <- APICallResult{URL: url, Error: fmt.Errorf("failed to create request: %v", err)}
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		results <- APICallResult{URL: url, Error: fmt.Errorf("failed to call service: %v", err)}
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		results <- APICallResult{URL: url, Error: fmt.Errorf("failed to read response body: %v", err)}
		return
	}

	results <- APICallResult{URL: url, Status: resp.Status, Response: string(body)}
	logger.WithFields(logrus.Fields{
		"url":    url,
		"status": resp.Status,
	}).Info("API call completed")
}

func main() {
	// Set log output to stdout and log level
	logger.SetOutput(os.Stdout)
	logger.SetLevel(logrus.InfoLevel)

	urls := []string{
		"https://jsonplaceholder.typicode.com/posts/1",
		"https://jsonplaceholder.typicode.com/posts/2",
		"https://jsonplaceholder.typicode.com/posts/3",
	}

	var wg sync.WaitGroup
	results := make(chan APICallResult, len(urls)) // Buffered channel for results

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	for _, url := range urls {
		wg.Add(1)
		go callThirdPartyService(ctx, url, &wg, results)
	}

	// Close the results channel after all Goroutines are done
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect results and handle errors
	for result := range results {
		if result.Error != nil {
			logger.WithFields(logrus.Fields{
				"url":   result.URL,
				"error": result.Error,
			}).Error("Error in API call")
		} else {
			fmt.Printf("Response from %s: %s\n", result.URL, result.Response)
		}
	}
}
