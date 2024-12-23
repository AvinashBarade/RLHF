package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// AlignedData represents a struct with fields optimized for memory alignment
type AlignedData struct {
	Value1 float64 // Aligned on 8-byte boundary
	Value2 float64 // Aligned on 8-byte boundary
}

// processBatchConcurrent uses Goroutines to process data in batches concurrently
func processBatchConcurrent(data []AlignedData, batchSize int, numWorkers int) {
	var wg sync.WaitGroup
	dataCh := make(chan []AlignedData, numWorkers) // Channel to distribute batches to workers

	// Start worker Goroutines
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for batch := range dataCh {
				for j := range batch {
					batch[j].Value1 *= 2
					batch[j].Value2 *= 3
				}
			}
		}()
	}

	// Distribute data batches to the channel
	for i := 0; i < len(data); i += batchSize {
		end := i + batchSize
		if end > len(data) {
			end = len(data)
		}
		dataCh <- data[i:end]
	}

	close(dataCh) // Close channel to signal workers
	wg.Wait()     // Wait for all workers to complete
}

func main() {
	// Large array of aligned data
	const dataSize = 1_000_000
	data := make([]AlignedData, dataSize)
	for i := range data {
		data[i] = AlignedData{Value1: float64(i), Value2: float64(i + 1)}
	}

	// Measure performance of regular range loop
	start := time.Now()
	processAligned(data)
	elapsed := time.Since(start)
	fmt.Printf("Regular range loop: %v\n", elapsed)

	// Measure performance of batch processing with Goroutines
	batchSize := 1024
	numWorkers := runtime.NumCPU()
	start = time.Now()
	processBatchConcurrent(data, batchSize, numWorkers)
	elapsed = time.Since(start)
	fmt.Printf("Batch processing with %d Goroutines: %v\n", numWorkers, elapsed)
}

// processAligned processes the data sequentially using a range loop
func processAligned(data []AlignedData) {
	for i := range data {
		data[i].Value1 *= 2
		data[i].Value2 *= 3
	}
}