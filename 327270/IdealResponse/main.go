package main

import (
	"testing"

	"github.com/benbjohnson/clock"
)

// Example function to benchmark
func exampleFunction(s string) []string {
	return []string{s, s + s}
}

// Benchmark function using an external library
func BenchmarkExampleFunction(b *testing.B) {
	// Set up clock for timing
	c := clock.New()

	// Setup input data for the benchmark
	s := "Hello, world!"

	// Time before execution
	start := c.Now()

	// Run the benchmark
	for i := 0; i < b.N; i++ {
		_ = exampleFunction(s)
	}

	// Log the time taken
	duration := c.Now().Sub(start)
	b.Logf("Time taken for %d iterations: %v", b.N, duration)
}

func main() {
	// This main function is just a placeholder since benchmarks are typically run via `go test`
}
