package main

import (
	"expvar"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

var (
	errorCount = expvar.NewInt("error_count")
)

func slowOperation() (int, error) {
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	if rand.Intn(2) == 0 {
		return 0, fmt.Errorf("random error occurred")
	}
	return 1, nil
}
func main() {
	rand.Seed(time.Now().UnixNano())

	go func() {
		// Export runtime metrics on HTTP port 8080
		if err := http.ListenAndServe(":8080", nil); err != nil {
			panic(err)
		}
	}()

	// Perform the slow operation and count the errors
	for {
		_, err := slowOperation()
		if err != nil {
			errorCount.Add(1)
		}
		time.Sleep(time.Second)
	}
}
