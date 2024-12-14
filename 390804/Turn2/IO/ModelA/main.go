package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

const numRequests = 5

func fetchURL(url string, wg *sync.WaitGroup) {
	defer wg.Done()
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("Fetched %s in %v with size %d bytes\n", url, time.Since(start), len(body))
}

func main() {
	var wg sync.WaitGroup
	urls := []string{
		"https://example.com",
		"https://golang.org",
		"https://httpbin.org/get",
		"https://jsonplaceholder.typicode.com/posts",
		"https://www.google.com",
	}

	wg.Add(numRequests)
	for _, url := range urls {
		go fetchURL(url, &wg)
	}

	wg.Wait()
}
