package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

func fetch(url string) (string, error) {
	// Fetch the URL
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return string(body), err
}

func main() {
	const numURLs = 5
	urls := []string{
		"http://example.com",
		"http://httpbin.org/get",
		"http://example.org",
		"http://golang.org",
		"https://jsonplaceholder.typicode.com/posts",
	}
	resultsCh := make(chan string)

	var wg sync.WaitGroup

	// Fixed-size WaitGroup to wait for all goroutines
	wg.Add(numURLs)

	// Goroutines for fetching URLs
	for _, url := range urls {
		go func(u string) {
			defer wg.Done() // Mark this goroutine as done when it finishes
			result, err := fetch(u)

			if err != nil {
				resultsCh <- fmt.Sprintf("Error fetching %s: %v", u, err)
			} else {
				resultsCh <- fmt.Sprintf("Fetched %s: %d characters", u, len(result))
			}
		}(url)
	}

	// Close the results channel when all fetching is done
	go func() {
		wg.Wait()
		close(resultsCh)
	}()

	// Collect results
	for result := range resultsCh {
		fmt.Println(result)
	}
}
