package main

import (
	"fmt"
	"time"
)

// Struct with aligned fields to optimize memory access
type AlignedData struct {
	Value1 float64 // Aligned on 8-byte boundary
	Value2 float64 // Aligned on 8-byte boundary
}

func processAligned(data []AlignedData) {
	for i := range data {
		data[i].Value1 *= 2
		data[i].Value2 *= 3
	}
}

// Batch processing for optimized iteration
func processBatch(data []AlignedData, batchSize int) {
	for i := 0; i < len(data); i += batchSize {
		end := i + batchSize
		if end > len(data) {
			end = len(data)
		}
		for j := i; j < end; j++ {
			data[j].Value1 *= 2
			data[j].Value2 *= 3
		}
	}
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

	// Measure performance of batch processing
	batchSize := 1024
	start = time.Now()
	processBatch(data, batchSize)
	elapsed = time.Since(start)
	fmt.Printf("Batch processing: %v\n", elapsed)
}
