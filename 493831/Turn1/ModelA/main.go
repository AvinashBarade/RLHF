package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
)

const (
	sliceLength   = 100000
	numGoroutines = 10
	iterations    = 1000000
)

func main() {
	// Initialize the shared slice
	slice := make([]int, sliceLength)
	for i := range slice {
		slice[i] = i
	}

	// Start goroutines to update the slice
	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go updateSlice(&wg, &slice)
	}

	wg.Wait()

	// Measure garbage collection statistics
	var gcStats runtime.GCStats
	runtime.ReadGCStats(&gcStats)
	fmt.Printf("GC Pauses: %d, Total GC Time: %dns\n", gcStats.NumGC, gcStats.PauseTotalNs)
}

func updateSlice(wg *sync.WaitGroup, slice *[]int) {
	defer wg.Done()

	for i := 0; i < iterations; i++ {
		// Randomly select an index and update the value
		index := rand.Intn(sliceLength)
		(*slice)[index] = rand.Int()
	}
}
