package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

const (
	numGoroutines    = 10
	workPerGoroutine = 1_000_000_000
)

var wg sync.WaitGroup

func cpuWork() {
	defer wg.Done()
	for i := 0; i < workPerGoroutine; i++ {
		for j := 0; j < 10000; j++ {
		}
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // Use all available CPU cores

	start := time.Now()

	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go cpuWork()
	}

	wg.Wait()

	end := time.Now()
	duration := end.Sub(start)
	fmt.Printf("Completed %d CPU-bound goroutines in %s\n", numGoroutines, duration)
}
