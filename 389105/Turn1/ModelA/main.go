package main

import (
	"fmt"
	"math"
	"time"

	"github.com/go-clockwork/timestamp"
)

func main() {
	// Creating two high precision timestamps
	now := timestamp.Now()
	oneHourAgo := now.Add(-1 * time.Hour)

	// Calculating the time interval between the two timestamps
	interval := now.Sub(oneHourAgo)

	// Converting the interval to milliseconds
	millis := float64(interval) / 1e6

	// Performing mathematical operations on the time interval
	// Example: Calculate square root of the interval in milliseconds
	sqrtMillis := math.Sqrt(millis)

	// Printing the results
	fmt.Printf("Current time: %s\n", now)
	fmt.Printf("Time one hour ago: %s\n", oneHourAgo)
	fmt.Printf("Time interval: %v\n", interval)
	fmt.Printf("Time interval in milliseconds: %.2f\n", millis)
	fmt.Printf("Square root of the interval in milliseconds: %.2f\n", sqrtMillis)
}
