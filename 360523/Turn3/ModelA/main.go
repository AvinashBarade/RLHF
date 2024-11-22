package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

const (
	numProducers = 10
	numConsumers = 5
	numTasks     = 10000
	bufferSize   = 100
)

var (
	wgUnbuffered sync.WaitGroup
	wgBuffered   sync.WaitGroup
	taskCounter  uint64 = 0
)

func produce(channel chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < numTasks; i++ {
		time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond) // Simulate work
		channel <- i
	}
}

func consume(channel <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for message := range channel {
		atomic.AddUint64(&taskCounter, 1) // Increment task counter for throughput measurement
		// Simulate work by performing some operation on the message
		time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
		fmt.Printf("Consumed message: %d\n", message)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	start := time.Now()

	// Unbuffered channel
	unbufferedChannel := make(chan int)

	// Start producers and consumers for unbuffered channel
	for i := 0; i < numProducers; i++ {
		wgUnbuffered.Add(1)
		go produce(unbufferedChannel, &wgUnbuffered)
	}
	for i := 0; i < numConsumers; i++ {
		wgUnbuffered.Add(1)
		go consume(unbufferedChannel, &wgUnbuffered)
	}

	// Buffered channel
	bufferedChannel := make(chan int, bufferSize)

	// Start producers and consumers for buffered channel
	for i := 0; i < numProducers; i++ {
		wgBuffered.Add(1)
		go produce(bufferedChannel, &wgBuffered)
	}
	for i := 0; i < numConsumers; i++ {
		wgBuffered.Add(1)
		go consume(bufferedChannel, &wgBuffered)
	}

	// Wait for all goroutines to finish
	wgUnbuffered.Wait()
	wgBuffered.Wait()

	end := time.Now()

	durationUnbuffered := end.Sub(start)
	fmt.Println("Unbuffered channel duration:", durationUnbuffered)
	fmt.Println("Unbuffered channel throughput:", taskCounter/uint64(durationUnbuffered.Seconds()))

	durationBuffered := end.Sub(start)
	fmt.Println("Buffered channel duration:", durationBuffered)
	fmt.Println("Buffered channel throughput:", taskCounter/uint64(durationBuffered.Seconds()))
}
