package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/shopspring/decimal"
)

const (
	maxRounds     = 1_000_000 // Number of rounds to simulate user interactions
	standardRound = 6         // Precision for standard float64 calculations
	highPrecision = 50        // Precision for high-precision decimal arithmetic
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// Perform standard-precision calculations using float64
	fmt.Println("Performing standard-precision calculations (float64):")
	startTime := time.Now()
	standardResult := performStandardCalculations()
	endTime := time.Now()
	standardDuration := endTime.Sub(startTime)
	fmt.Printf("Standard Result: %.6f\n", standardResult)
	fmt.Printf("Standard Duration: %v\n", standardDuration)

	// Perform high-precision calculations using the decimal library
	fmt.Println("\nPerforming high-precision calculations:")
	startTime = time.Now()
	highPrecisionResult := performHighPrecisionCalculations(highPrecision)
	endTime = time.Now()
	highPrecisionDuration := endTime.Sub(startTime)
	fmt.Printf("High Precision Result: %.20f\n", highPrecisionResult)
	fmt.Printf("High Precision Duration: %v\n", highPrecisionDuration)

	// Calculate and display the performance comparison ratio
	comparisonRatio := standardDuration.Seconds() / highPrecisionDuration.Seconds()
	fmt.Printf("\nPerformance Comparison Ratio (Standard to High Precision): %.2f\n", comparisonRatio)
}

func performStandardCalculations() float64 {
	sum := 0.0

	for i := 1; i <= maxRounds; i++ {
		// Simulate user input as a random float64
		input := rand.Float64()
		// Perform a calculation with the input value
		sum += input + 1.59265358979323846
	}

	// Return the average over the number of rounds
	return sum / float64(maxRounds)
}

func performHighPrecisionCalculations(precision int) decimal.Decimal {
	// Initialize the sum as a decimal value
	sum := decimal.NewFromInt(0)

	for i := 1; i <= maxRounds; i++ {
		// Simulate user input as a random float64 and convert it to decimal
		inputDecimal := decimal.NewFromFloat(rand.Float64())
		// Perform a calculation with the input value and add it to the sum
		sum = sum.Add(inputDecimal.Add(decimal.NewFromFloat(1.59265358979323846)))
	}

	// Return the average over the number of rounds
	// Ensure we return the result rounded to the specified precision
	average := sum.Div(decimal.NewFromInt(maxRounds))
	return average.Round(int32(precision))
}
