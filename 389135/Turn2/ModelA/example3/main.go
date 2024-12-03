package main

import (
	"fmt"
	"sync"
)

func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func calculateFibonacci(start, end int, results chan<- int) {
	var sum int
	for i := start; i <= end; i++ {
		sum += fibonacci(i)
	}
	results <- sum
}

func main() {
	const numGoroutines = 4
	const numIterations = 1000000

	var results chan<- int = make(chan int)
	var wg sync.WaitGroup

	chunkSize := numIterations / numGoroutines

	for i := 0; i < numGoroutines; i++ {
		start := i * chunkSize
		end := min((i+1)*chunkSize-1, numIterations-1)

		wg.Add(1)
		go func(start, end int) {
			calculateFibonacci(start, end, results)
			wg.Done()
		}(start, end)
	}

	go func() {
		var total int
		for result := range results {
			total += result
		}
		fmt.Println("Total sum of Fibonacci numbers:", total)
		close(results)
	}()

	wg.Wait()
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
