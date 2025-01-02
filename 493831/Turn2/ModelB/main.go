package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

const (
	maxSliceLength = 100000 // Maximum slice size handled by the pool
	numGoroutines  = 10
	iterations     = 10000
)

type slicePool struct {
	pools []*sync.Pool // Slice of memory pools for different slice sizes
}

func newSlicePool() *slicePool {
	p := &slicePool{
		pools: make([]*sync.Pool, maxSliceLength+1),
	}

	// Initialize pools for slice sizes from 1 to maxSliceLength
	for i := range p.pools {
		p.pools[i] = &sync.Pool{
			New: func() interface{} {
				return make([]int, i)
			},
		}
	}
	return p
}

func (p *slicePool) GetSlice(length int) []int {
	// Ensure length is within the bounds of the pool
	if length < 1 || length > maxSliceLength {
		panic("Invalid slice length requested")
	}
	return p.pools[length].Get().([]int)
}

func (p *slicePool) PutSlice(slice []int) {
	// Put the slice back into the appropriate pool
	p.pools[cap(slice)].Put(slice)
}

func main() {
	pool := newSlicePool()

	// Track start time
	start := time.Now()

	// Start goroutines to update the slice
	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go updateSlice(&wg, pool)
	}

	wg.Wait()

	// Measure elapsed time
	elapsed := time.Since(start)

	// Collect garbage collection statistics
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	fmt.Printf("Execution Time: %v\n", elapsed)
	fmt.Printf("GC Count: %d, Total Alloc: %d KB, Total GC Pause Time: %d ns\n",
		memStats.NumGC, memStats.TotalAlloc/1024, memStats.PauseTotalNs)
}

func updateSlice(wg *sync.WaitGroup, pool *slicePool) {
	defer wg.Done()

	for i := 0; i < iterations; i++ {
		// Randomly determine the slice size
		sliceSize := rand.Intn(maxSliceLength) + 1

		// Get a slice from the pool
		slice := pool.GetSlice(sliceSize)

		// Update the slice elements
		for j := range slice {
			slice[j] = rand.Int()
		}

		// Put the slice back into the pool
		pool.PutSlice(slice)
	}
}
