package main

import (
	"fmt"
	"log"
	"math/rand/v2"
	"sync"
	"time"
)

func crawl(urls <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for url := range urls {
		log.Printf("Fetching: %s\n", url)
		time.Sleep(time.Duration(int64(rand.Float64()*1000)) * time.Millisecond)
		log.Printf("Processed: %s\n", url)
	}
}

func main() {
	const numProducers = 3
	const numConsumers = 3
	const numURLs = 100

	urls := make(chan string, numURLs*2) // Buffered channel to manage queue
	wg := &sync.WaitGroup{}

	// Start consumers
	for i := 0; i < numConsumers; i++ {
		wg.Add(1)
		go crawl(urls, wg)
	}

	// Start producers
	for i := 0; i < numURLs; i++ {
		url := fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/%d", i+1)
		urls <- url
		log.Printf("Added URL: %s\n", url)
	}

	// Close the channel to signal consumers to stop
	close(urls)
	wg.Wait()
}
