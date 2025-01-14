package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	numOps      = flag.Int("num_ops", 100000, "number of operations")
	concurrency = flag.Int("concurrency", 1, "number of goroutines")
)

func main() {
	flag.Parse()
	rand.Seed(time.Now().UnixNano())

	benchmark("Insert")
	benchmark("Lookup")
	benchmark("Delete")
}

func benchmark(operation string) {
	start := time.Now()

	var wg sync.WaitGroup
	m := make(map[int]int)

	switch operation {
	case "Insert":
		benchmarkInsert(&wg, m)
	case "Lookup":
		benchmarkLookup(&wg, m)
	case "Delete":
		benchmarkDelete(&wg, m)
	default:
		fmt.Println("Invalid operation")
		return
	}

	wg.Wait()

	duration := time.Since(start)
	fmt.Printf("Benchmark %s completed in %v with %d goroutines and %d operations\n", operation, duration, *concurrency, *numOps)
}

func benchmarkInsert(wg *sync.WaitGroup, m map[int]int) {
	defer wg.Done()
	for i := 0; i < *numOps; i++ {
		key := rand.Intn(10000)
		m[key] = key
	}
}

func benchmarkLookup(wg *sync.WaitGroup, m map[int]int) {
	defer wg.Done()
	for i := 0; i < *numOps; i++ {
		key := rand.Intn(10000)
		_, _ = m[key]
	}
}

func benchmarkDelete(wg *sync.WaitGroup, m map[int]int) {
	defer wg.Done()
	for i := 0; i < *numOps; i++ {
		key := rand.Intn(10000)
		delete(m, key)
	}
}
