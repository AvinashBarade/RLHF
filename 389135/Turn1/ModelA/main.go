package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const numGoroutines = 5
const numIterations = 1000000

func randomTask(wg *sync.WaitGroup, num int) {
	defer wg.Done()
	for i := 0; i < num; i++ {
		_ = rand.Intn(10000)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var wg sync.WaitGroup

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go randomTask(&wg, numIterations/numGoroutines)
	}

	wg.Wait()
	fmt.Println("Done.")
}
