package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	numSlices := 10000
	sliceSize := 1000
	slices := make([]int, numSlicessliceSize)
	var wg sync.WaitGroup
	wg.Add(numSlices)
	start := time.Now()
	for i := 0; i < numSlices; i++ {
		go updateSlice(isliceSize, (i+1)*sliceSize, slices, &wg)
	}
	wg.Wait()
	end := time.Now()
	fmt.Printf("Time taken: %v\n", end.Sub(start))
}

func updateSlice(start int, end int, slices []int, wg *sync.WaitGroup) {
	for i := start; i < end; i++ {
		slices[i] = rand.Int()
	}
	wg.Done()
}
