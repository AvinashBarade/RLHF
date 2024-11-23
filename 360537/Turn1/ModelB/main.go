package main

import (
	"math/rand"
	"sync"
	"testing"
)

const (
	numOps        = 1000000
	numGoroutines = 10
)

var (
	wg sync.WaitGroup
)

func TestMapPerformance(t *testing.T) {
	m := make(map[int]int)

	b.Run("MapReads", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			readFromMap(m)
		}
	})

	b.Run("MapWrites", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			writeToMap(m)
		}
	})
}

func TestSlicePerformance(t *testing.T) {
	s := make([]int, numOps)

	b.Run("SliceReads", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			readFromSlice(s)
		}
	})

	b.Run("SliceWrites", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			writeToSlice(s)
		}
	})
}

func readFromMap(m map[int]int) {
	for i := 0; i < numOps; i++ {
		_, _ = m[rand.Intn(numOps)]
	}
}

func writeToMap(m map[int]int) {
	for i := 0; i < numOps; i++ {
		m[i] = i
	}
}

func readFromSlice(s []int) {
	for i := 0; i < numOps; i++ {
		_ = s[rand.Intn(numOps)]
	}
}

func writeToSlice(s []int) {
	for i := 0; i < numOps; i++ {
		s[i] = i
	}
}
