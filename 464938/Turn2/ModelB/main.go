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

// Struct with unaligned fields to demonstrate the impact
type UnalignedData struct {
	Value1 float64
	Value2 float64
}

func processAligned(data []AlignedData) {
	for i := range data {
		data[i].Value1 *= 2
		data[i].Value2 *= 3
	}
}

func processUnaligned(data []UnalignedData) {
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
	const dataSize = 1_000_000

	// Large array of aligned data
	alignedData := make([]AlignedData, dataSize)
	for i := range alignedData {
		alignedData[i] = AlignedData{Value1: float64(i), Value2: float64(i + 1)}
	}

	// Large array of unaligned data
	unalignedData := make([]UnalignedData, dataSize)
	for i := range unalignedData {
		unalignedData[i] = UnalignedData{Value1: float64(i), Value2: float64(i + 1)}
	}

	// Measure performance of regular range loop on aligned data
	start := time.Now()
	processAligned(alignedData)
	elapsedAligned := time.Since(start)
	fmt.Printf("Regular range loop on aligned data: %v\n", elapsedAligned)

	// Measure performance of regular range loop on unaligned data
	start = time.Now()
	processUnaligned(unalignedData)
	elapsedUnaligned := time.Since(start)
	fmt.Printf("Regular range loop on unaligned data: %v\n", elapsedUnaligned)

	// Measure performance of batch processing on aligned data
	batchSize := 1024
	start = time.Now()
	processBatch(alignedData, batchSize)
	elapsedBatchAligned := time.Since(start)
	fmt.Printf("Batch processing on aligned data: %v\n", elapsedBatchAligned)

	// Measure performance of batch processing on unaligned data
	start = time.Now()
	processBatch(alignedData, batchSize)
	elapsedBatchUnaligned := time.Since(start)
	fmt.Printf("Batch processing on unaligned data: %v\n", elapsedBatchUnaligned)
}
