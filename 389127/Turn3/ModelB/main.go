package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/shopspring/decimal"
)

const (
	maxRounds     = 1000000
	standardRound = 6   // For float64
	highPrecision = 256 // For decimal.Decimal
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// Perform standard precision calculations
	fmt.Println("\nStarting standard precision calculations:")
	startTimeStandard := time.Now()
	resultStandard := performCalculations(standardRound)
	endTimeStandard := time.Now()
	standardDuration := endTimeStandard.Sub(startTimeStandard)
	fmt.Printf("Standard Precision Result: %.6f\n", resultStandard)
	fmt.Printf("Standard Precision Duration: %v\n", standardDuration)

	// Perform high precision calculations
	fmt.Println("\nStarting high precision calculations:")
	startTimeHigh := time.Now()
	resultHigh := performCalculations(highPrecision)
	endTimeHigh := time.Now()
	highPrecisionDuration := endTimeHigh.Sub(startTimeHigh)
	fmt.Printf("High Precision Result: %s\n", resultHigh)
	fmt.Printf("High Precision Duration: %v\n", highPrecisionDuration)

	// Calculate and display the performance comparison ratio
	comparisonRatio := float64(standardDuration) / float64(highPrecisionDuration)
	fmt.Printf("Performance Comparison Ratio (Standard to High Precision): %.2f\n", comparisonRatio)
}

func performCalculations(precision int) interface{} {
	sum := decimal.Zero
	var result interface{}

	if precision == standardRound {
		for i := 1; i <= maxRounds; i++ {
			// Simulate user input as a random float
			input := rand.Float64()
			// Perform a calculation with the input value
			sum += input + 1.59265358979323846
		}
		// Convert sum back to float64 for standard precision result
		result = sum / float64(maxRounds)
	} else {
		for i := 1; i <= maxRounds; i++ {
			// Simulate user input as a random float converted to decimal.Decimal
			input := decimal.NewFromFloat(rand.Float64())
			// Perform a calculation with the input value using decimal.Decimal
			sum = sum.Add(input.Add(decimal.NewFromFloat(1.59265358979323846)))
		}
		// Calculate average and round to high precision
		result = sum.Div(decimal.NewFromInt(maxRounds)).Round(precision)
	}
	return result
}
