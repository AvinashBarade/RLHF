package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

const (
	numProducers = 10
	numConsumers = 5
	numTasks     = 1_000_000
	bufferSize   = 1000
)

var (
	wg             sync.WaitGroup
	totalTasks     int32
	processedTasks int32
)

func produce(channel chan<- int) {
	defer wg.Done()
	for i := 0; i < numTasks/numProducers; i++ {
		time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond) // Simulate work
		channel <- i
		atomic.AddInt32(&totalTasks, 1)
	}
}

func consume(channel <-chan int) {
	defer wg.Done()
	for {
		select {
		case message := <-channel:
			time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond) // Simulate work
			atomic.AddInt32(&processedTasks, 1)
		default:
			// No more messages to process, exit
			if atomic.LoadInt32(&processedTasks) == atomic.LoadInt32(&totalTasks) {
				return
			}
			// Yield to other goroutines to avoid CPU hogging
			runtime.Gosched()
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	start := time.Now()

	unbufferedChannel := make(chan int)
	bufferedChannel := make(chan int, bufferSize)

	// Start producers and consumers for unbuffered channel
	wg.Add(numProducers)
	for i := 0; i < numProducers; i++ {
		go produce(unbufferedChannel)
	}
	wg.Add(numConsumers)
	for i := 0; i < numConsumers; i++ {
		go consume(unbufferedChannel)
	}

	// Start producers and consumers for buffered channel
	wg.Add(numProducers)
	for i := 0; i < numProducers; i++ {
		go produce(bufferedChannel)
	}
	wg.Add(numConsumers)
	for i := 0; i < numConsumers; i++ {
		go consume(bufferedChannel)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	end := time.Now()

	durationUnbuffered := end.Sub(start)
	fmt.Println("Unbuffered channel duration:", durationUnbuffered)

	start = time.Now()
	// Wait for all goroutines to finish with buffered channel
	wg.Wait()
	end = time.Now()

	durationBuffered := end.Sub(start)
	fmt.Println("Buffered channel duration:", durationBuffered)

	// Performance comparison
	fmt.Println("Throughput (Unbuffered):", float64(numTasks)/durationUnbuffered.Seconds())
	fmt.Println("Throughput (Buffered):", float64(numTasks)/durationBuffered.Seconds())
	fmt.Println("Latency (Unbuffered):", durationUnbuffered/time.Duration(numTasks))
	fmt.Println("Latency (Buffered):", durationBuffered/time.Duration(numTasks))
}
