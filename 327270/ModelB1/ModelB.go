package main

import (
	"fmt"

	"github.com/prataprc/gobenchmarks"
)

// Factorial function
func Factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * Factorial(n-1)
}

// BenchmarkFactorial function to benchmark the Factorial function
func BenchmarkFactorial(b *gobenchmarks.B) {
	for i := 0; i < b.N; i++ {
		Factorial(20)
	}
}
func main() {
	fmt.Println("Factorial(20) =", Factorial(20))
	gobenchmarks.Run("BenchmarkFactorial")
}
