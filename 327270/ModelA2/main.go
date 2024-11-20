package main

import (
	"testing"

	"go.opentelemetry.io/contrib/instrumentation/go/benchmark"
)

func exampleFunction(s string) []string {
	return []string{s, s + s}
}

func BenchmarkExampleFunction(b *testing.B) {
	s := "Hello, world!"
	for i := 0; i < b.N; i++ {
		_ = exampleFunction(s)
	}
}

func BenchmarkExampleFunctionWithOTel(b *benchmark.Benchmarker) {
	s := "Hello, world!"
	for i := 0; i < b.N; i++ {
		_ = exampleFunction(s)
	}
}

func TestMain(m *testing.M) {
	// Initialize OpenTelemetry benchmarking framework
	benchmark.SetupBenchmarking(m)
}
