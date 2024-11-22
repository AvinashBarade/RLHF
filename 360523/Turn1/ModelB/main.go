package main

import (
	"fmt"
	"net/http"
	"time"
)

func slowComputation(w http.ResponseWriter, r *http.Request, duration time.Duration) {
	// Simulate a slow computation by sleeping for the specified duration
	time.Sleep(duration)
	fmt.Fprintf(w, "Slow computation complete.\n")
}

func startServer(reqCh <-chan *http.Request, duration time.Duration) {
	for req := range reqCh {
		slowComputation(nil, req, duration)
	}
}
func main() {
	numWorkers := 10
	buffered := 1000

	// Create an unbuffered channel
	unbufferedReqCh := make(chan *http.Request)

	// Create a buffered channel
	bufferedReqCh := make(chan *http.Request, buffered)

	// Start the worker goroutines
	for i := 0; i < numWorkers; i++ {
		go startServer(unbufferedReqCh, 1*time.Second) // Unbuffered channel
		go startServer(bufferedReqCh, 1*time.Second)   // Buffered channel
	}

	// Start the web server to handle incoming requests
	http.HandleFunc("/unbuffered", func(w http.ResponseWriter, r *http.Request) {
		unbufferedReqCh <- r
	})
	http.HandleFunc("/buffered", func(w http.ResponseWriter, r *http.Request) {
		bufferedReqCh <- r
	})

	fmt.Println("Server running on http://localhost:8080")
	fmt.Println("Press Ctrl+C to exit.")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
