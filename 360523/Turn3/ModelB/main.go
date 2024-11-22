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
		time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond) // Simulate work
		processedTaskCount := atomic.AddUint64(&taskCounter, 1)
		fmt.Printf("Consumed task: %d (Total processed: %d)\n", message, processedTaskCount)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	start := time.Now()

	unbufferedChannel := make(chan int)
	bufferedChannel := make(chan int, bufferSize)
	for i := 0; i < numProducers; i++ {
		wgUnbuffered.Add(1)
		go produce(unbufferedChannel, &wgUnbuffered)
		wgBuffered.Add(1)
		go produce(bufferedChannel, &wgBuffered)
	}

	for i := 0; i < numConsumers; i++ {
		wgUnbuffered.Add(1)
		go consume(unbufferedChannel, &wgUnbuffered)
		wgBuffered.Add(1)
		go consume(bufferedChannel, &wgBuffered)
	}
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
