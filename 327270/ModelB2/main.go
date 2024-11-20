package main

import (
	"fmt"
	"testing"

	"github.com/google/gobenchs"
)

func exampleFunction(s string) []string {
	return []string{s, s + s}
}

func BenchmarkExampleFunction(b *testing.B) {
	for i := 0; i < b.N; i++ {
		exampleFunction("test")
	}
}

func main() {
	res, err := gobenchs.Run("-format csv")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	// Process and display the benchmark results using the CSV format
	fmt.Println(res.String())
}
