package main

import (
	"testing"
	"time"

	"github.com/benbjohnson/clock"
)

func TestTimeSensitiveFunction(t *testing.T) {
	realClock := clock.NewRealClock()
	mockClock := clock.NewMockClock()

	// Set initial time
	mockClock.Set(realClock.Now())

	// Test without delay
	if !TimeSensitiveFunction(mockClock, 0) {
		t.Error("TimeSensitiveFunction failed without delay")
	}

	// Test with delay
	mockClock.Set(mockClock.Now().Add(time.Second)) // Move time ahead
	if TimeSensitiveFunction(mockClock, 1*time.Second) {
		t.Error("TimeSensitiveFunction passed with delay")
	}
}
