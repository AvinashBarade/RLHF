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
	fmt.Printf("High Precision Result: %.20f\n", highPrecisionResult.Round(standardRound))
	fmt.Printf("High Precision Duration: %v\n", highPrecisionDuration)

	// Calculate and display the performance comparison ratio
	comparisonRatio := standardDuration.Seconds() / highPrecisionDuration.Seconds()
	fmt.Printf("\nPerformance Comparison Ratio (Standard to High Precision): %.2f\n", comparisonRatio)
}

func performStandardCalculations() float64 {
	sum := 0.0

	for i := 1; i <= maxRounds; i++ {
		// Simulate user input as a random float64
		input := float64(rand.Float64())
		// Perform a calculation with the input value
		sum += input + 1.59265358979323846
	}

	// Return the average over the number of rounds
	return sum / float64(maxRounds)
}

func performHighPrecisionCalculations(precision int) decimal.Decimal {
	precisionCtx := decimal.NewContext(precision, 2)
	sum := decimal.NewFromInt(0)

	for i := 1; i <= maxRounds; i++ {
		// Simulate user input as a random float64 and convert it to decimal
		inputDecimal := decimal.NewFromFloat(float64(rand.Float64())).SetScale(10) // Adjust scale for decimal
		// Perform a calculation with the input value
		sum = precisionCtx.Add(sum, precisionCtx.Add(inputDecimal, decimal.NewFromFloat(1.59265358979323846)))
	}

	// Return the average over the number of rounds, using standardized scale
	return precisionCtx.Div(sum, decimal.NewFromInt(maxRounds))
}
