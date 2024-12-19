package main

import (
	"fmt"
	"runtime"
	"time"
)

// Struct with aligned fields to optimize memory access
type AlignedData struct {
	Value1 float64 // Aligned on 8-byte boundary
	Value2 float64 // Aligned on 8-byte boundary
}

// Struct with unaligned fields
type UnalignedData struct {
	Value1 float64
	Value2 int32 // Unaligned, causing padding issues
	Value3 float64
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
		data[i].Value2 += 42
		data[i].Value3 *= 5
	}
}

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
	runtime.GOMAXPROCS(1) // Single thread for consistent timing

	// Large array of aligned data
	const dataSize = 1_000_000
	alignedData := make([]AlignedData, dataSize)
	for i := range alignedData {
		alignedData[i] = AlignedData{Value1: float64(i), Value2: float64(i + 1)}
	}

	// Large array of unaligned data
	unalignedData := make([]UnalignedData, dataSize)
	for i := range unalignedData {
		unalignedData[i] = UnalignedData{Value1: float64(i), Value2: int32(i + 1), Value3: float64(i + 2)}
	}

	// Measure performance of regular range loop with aligned data
	start := time.Now()
	processAligned(alignedData)
	elapsedAligned := time.Since(start)
	fmt.Printf("Regular range loop with aligned data: %v\n", elapsedAligned)

	// Measure performance of regular range loop with unaligned data
	start = time.Now()
	processUnaligned(unalignedData)
	elapsedUnaligned := time.Since(start)
	fmt.Printf("Regular range loop with unaligned data: %v\n", elapsedUnaligned)

	// Measure performance of batch processing with aligned data
	batchSize := 1024
	start = time.Now()
	processBatch(alignedData, batchSize)
	elapsedBatchAligned := time.Since(start)
	fmt.Printf("Batch processing with aligned data: %v\n", elapsedBatchAligned)
}
