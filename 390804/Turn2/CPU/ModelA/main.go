package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

const (
	numGoroutines    = 8
	workPerGoroutine = 100000000 // Number of iterations in the busy loop
)

func computeFibonacci(n int, wg *sync.WaitGroup) {
	defer wg.Done()
	a, b := 0, 1
	for i := 0; i < n; i++ {
		a, b = b, a+b
	}
	// Just to prevent compiler optimization
	_ = a
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // Use all available CPU cores

	var wg sync.WaitGroup
	start := time.Now()

	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go computeFibonacci(workPerGoroutine, &wg)
	}

	wg.Wait()
	duration := time.Since(start)
	fmt.Printf("Completed %d goroutines in %s\n", numGoroutines, duration)
}
