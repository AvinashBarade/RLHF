package main

import (
	"testing"

	"github.com/google/go-globmark"
)

func exampleFunction(s string) []string {
	return []string{s, s + s}
}

func init() {
	globmark.RegisterBenchmarkString("ExampleFunction", exampleFunction)
}

func TestExampleFunction(t *testing.T) {
	globmark.TestBenchmarks(t)
}
