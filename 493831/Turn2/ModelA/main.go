package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

const (
	sliceLength   = 100000
	numGoroutines = 10
	iterations    = 10000
	maxPoolSize   = 100 // Maximum number of slices to keep in each pool
)

// DynamicSlicePool manages slices of varying sizes dynamically
type DynamicSlicePool struct {
	pools map[int]*sync.Pool
}

func (dsp *DynamicSlicePool) Get(size int) []int {
	if pool, ok := dsp.pools[size]; ok {
		if item := pool.Get(); item != nil {
			return item.([]int)
		}
	}
	return make([]int, size)
}

func (dsp *DynamicSlicePool) Put(size int, slice []int) {
	if len(slice) != size {
		return
	}

	if pool, ok := dsp.pools[size]; ok {
		if atomic.LoadInt32(&pool.Len()) < maxPoolSize {
			pool.Put(slice)
		}
	} else {
		newPool := &sync.Pool{New: func() interface{} { return make([]int, size) }}
		dsp.pools[size] = newPool
		newPool.Put(slice)
	}
}

func main() {
	// Initialize the shared slice pool
	dsp := &DynamicSlicePool{pools: make(map[int]*sync.Pool)}

	// Track start time
	start := time.Now()

	// Start goroutines to update the slice
	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go updateSlice(&wg, dsp)
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

func updateSlice(wg *sync.WaitGroup, dsp *DynamicSlicePool) {
	defer wg.Done()

	for i := 0; i < iterations; i++ {
		// Randomly select a slice size and update the slice
		size := rand.Intn(sliceLength/10) + 1 // Random size between 1 and sliceLength/10
		slice := dsp.Get(size)

		for j := 0; j < size; j++ {
			slice[j] = rand.Int()
		}

		dsp.Put(size, slice)
	}
}
