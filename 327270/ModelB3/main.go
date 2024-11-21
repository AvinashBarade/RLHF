package main

import (
	"testing"
	"time"

	"github.com/benbjohnson/clock"
)

func exampleFunction(s string, clock clock.Clock) []string {
	// Use the provided clock for time-related operations
	now := clock.Now()
	// ...
	// Time-based operations
	// ...
	return []string{s, s + s}
}

func BenchmarkExampleFunction(b *testing.B) {
	sysClock := clock.NewSystem()
	b.ResetTimer()

	s := "Hello, world!"
	for i := 0; i < b.N; i++ {
		_ = exampleFunction(s, sysClock)
	}
}
func BenchmarkExampleFunctionWithMockClock(b *testing.B) {
	mockClock := clock.NewMock()
	// Manually control the passage of time
	ticker := time.NewTicker(500 * time.Millisecond)
	go func() {
		for range ticker.C {
			mockClock.Add(time.Second)
		}
	}()
	defer ticker.Stop()

	b.ResetTimer()

	s := "Hello, world!"
	for i := 0; i < b.N; i++ {
		_ = exampleFunction(s, mockClock)
	}
}
