package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var wg sync.WaitGroup

func fetchURL(url string) {
	defer wg.Done()
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching URL:", err)
		return
	}
	defer resp.Body.Close()
	fmt.Println("Fetched:", url)
}

func main() {
	urls := []string{
		"https://example.com",
		"https://google.com",
		"https://github.com",
		// Add more URLs as needed
	}

	start := time.Now()

	wg.Add(len(urls))
	for _, url := range urls {
		go fetchURL(url)
	}

	wg.Wait()

	end := time.Now()
	duration := end.Sub(start)
	fmt.Printf("Fetched all URLs in %s\n", duration)
}
