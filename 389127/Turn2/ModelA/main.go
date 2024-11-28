package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/shopspring/decimal"
)

const (
	maxRounds     = 1000000
	standardRound = 6
	highPrecision = 256
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// Perform standard-precision calculations
	startTime := time.Now()
	standardResult := performCalculations(standardRound)
	endTime := time.Now()
	standardDuration := endTime.Sub(startTime)

	// Perform high-precision calculations
	startTime = time.Now()
	highPrecisionResult := performCalculations(highPrecision)
	endTime = time.Now()
	highPrecisionDuration := endTime.Sub(startTime)

	// Display results and analysis
	fmt.Printf("Standard Precision Result (after %d rounds): %s\n", maxRounds, standardResult)
	fmt.Printf("Standard Precision Duration: %v\n", standardDuration)

	fmt.Printf("High Precision Result (after %d rounds): %s\n", maxRounds, highPrecisionResult)
	fmt.Printf("High Precision Duration: %v\n", highPrecisionDuration)

	// Calculate and display the performance comparison ratio
	comparisonRatio := float64(standardDuration) / float64(highPrecisionDuration)
	fmt.Printf("Performance Comparison Ratio (Standard to High Precision): %.2f\n", comparisonRatio)
}

func performCalculations(precision int) decimal.Decimal {
	sum := decimal.Zero

	for i := 1; i <= maxRounds; i++ {
		// Simulate user input as a random float
		input := decimal.NewFromFloat(rand.Float64())

		// Perform a calculation with the input value
		sum = sum.Add(input.Add(decimal.NewFromFloat(1.59265358979323846)))
	}

	// Return the sum rounded to the specified precision
	return sum.Div(decimal.NewFromInt(maxRounds)).Round(precision)
}
