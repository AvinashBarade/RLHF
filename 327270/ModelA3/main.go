package main

import (
	"time"

	"github.com/benbjohnson/clock"
)

// TimeSensitiveFunction introduces a delay before executing the main task.
func TimeSensitiveFunction(c clock.Clock, delay time.Duration) bool {
	t := c.Now()
	if delay > 0 {
		c.Sleep(delay)
	}
	// Simulate some work here
	// For simplicity, just check if the delay was longer than 500ms
	return t.Add(500 * time.Millisecond).Before(c.Now())
}
