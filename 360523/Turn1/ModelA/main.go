package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numProducers = 10
	numConsumers = 5
	numTasks     = 10000
	bufferSize   = 100
)

var (
	wg sync.WaitGroup
)

func produce(channel chan<- int) {
	defer wg.Done()
	for i := 0; i < numTasks; i++ {
		time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond) // Simulate work
		channel <- i
	}
}

func consume(channel <-chan int) {
	defer wg.Done()
	for message := range channel {
		time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond) // Simulate work
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	start := time.Now()

	unbufferedChannel := make(chan int)
	bufferedChannel := make(chan int, bufferSize)

	// Start producers and consumers for unbuffered channel
	for i := 0; i < numProducers; i++ {
		wg.Add(1)
		go produce(unbufferedChannel)
	}
	for i := 0; i < numConsumers; i++ {
		wg.Add(1)
		go consume(unbufferedChannel)
	}

	// Start producers and consumers for buffered channel
	for i := 0; i < numProducers; i++ {
		wg.Add(1)
		go produce(bufferedChannel)
	}
	for i := 0; i < numConsumers; i++ {
		wg.Add(1)
		go consume(bufferedChannel)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	end := time.Now()

	durationUnbuffered := end.Sub(start)
	fmt.Println("Unbuffered channel duration:", durationUnbuffered)
	durationBuffered := end.Sub(start)
	fmt.Println("Buffered channel duration:", durationBuffered)
}
