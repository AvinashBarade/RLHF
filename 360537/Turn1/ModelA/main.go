package main

import (
	"fmt"
	"runtime"
	"testing"
)

var n = 100000

func BenchmarkMapRead(b *testing.B) {
	var m map[int]int
	for i := 0; i < n; i++ {
		m[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m[i%n] // Just accessing the element
	}
}

func BenchmarkMapWrite(b *testing.B) {
	var m map[int]int

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m[i%n] = i % n
	}
}

func BenchmarkSliceRead(b *testing.B) {
	var s []int
	for i := 0; i < n; i++ {
		s = append(s, i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s[i%n] // Just accessing the element
	}
}

func BenchmarkSliceWrite(b *testing.B) {
	var s []int

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s = append(s, i)
	}
}

func main() {
	fmt.Printf("Go version: %s, architecture: %s\n", runtime.Version(), runtime.GOARCH)

	testing.Run(nil, func(t *testing.T) {
		testing.Benchmark(BenchmarkMapRead)
		testing.Benchmark(BenchmarkMapWrite)
		testing.Benchmark(BenchmarkSliceRead)
		testing.Benchmark(BenchmarkSliceWrite)
	})
}
