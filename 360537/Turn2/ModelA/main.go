package main

import (
	"fmt"
	"runtime"
	"testing"
)

func BenchmarkMapRead(b *testing.B) {
	// Initialize a map with n keys and values
	m := make(map[int]int, n)
	for i := 0; i < n; i++ {
		m[i] = i
	}

	// Perform the read operation b.N times
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = m[i%n] // Read the value at the index (i%n)
	}
}

func BenchmarkMapWrite(b *testing.B) {
	// Initialize an empty map
	m := make(map[int]int)

	// Perform the write operation b.N times
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m[i%n] = i % n // Write the value i%n at the index (i%n)
	}
}

func BenchmarkSliceRead(b *testing.B) {
	// Initialize a slice with n elements
	s := make([]int, n)
	for i := 0; i < n; i++ {
		s[i] = i
	}

	// Perform the read operation b.N times
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = s[i%n] // Read the value at the index (i%n)
	}
}

func BenchmarkSliceWrite(b *testing.B) {
	// Initialize an empty slice
	s := make([]int, 0, n) // Preallocate space for n elements

	// Perform the write operation b.N times
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s = append(s, i%n) // Append the value i%n to the slice
	}
}

func main() {
	fmt.Printf("Go version: %s, architecture: %s\n", runtime.Version(), runtime.GOARCH)

	// Run the benchmarks
	testing.Benchmark(BenchmarkMapRead)
	testing.Benchmark(BenchmarkMapWrite)
	testing.Benchmark(BenchmarkSliceRead)
	testing.Benchmark(BenchmarkSliceWrite)
}
