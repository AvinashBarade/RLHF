package main

import (
	"fmt"
	"runtime"
	"testing"
)

var n = 100000 // Number of elements in the datasets

// This function initializes a map with n elements.
func createMap() map[int]int {
	m := make(map[int]int)
	for i := 0; i < n; i++ {
		m[i] = i
	}
	return m
}

// This function initializes a slice with n elements.
func createSlice() []int {
	s := make([]int, n)
	for i := 0; i < n; i++ {
		s[i] = i
	}
	return s
}

func BenchmarkMapRead(b *testing.B) {
	m := createMap() // Initialize map once for each benchmark run

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = m[i%n] // _ is used to ignore the unused result of m[i%n]
	}
}

func BenchmarkMapWrite(b *testing.B) {
	m := createMap()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m[i%n] = i % n
	}
}

func BenchmarkSliceRead(b *testing.B) {
	s := createSlice()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = s[i%n]
	}
}

func BenchmarkSliceWrite(b *testing.B) {
	s := createSlice()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s[i%n] = i % n
	}
}

func main() {
	fmt.Printf("Go version: %s, architecture: %s\n", runtime.Version(), runtime.GOARCH)
	testing.Main(nil, nil, nil, nil) // Run all tests and benchmarks automatically
}
