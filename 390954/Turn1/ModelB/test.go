package main

import (
	"fmt"
	"sync"
	"testing"
)

// using sync.Mutex
func BenchmarkMutex(b *testing.B) {
	var counter int
	var mutex sync.Mutex

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			mutex.Lock()
			counter++
			mutex.Unlock()
		}
	})
	fmt.Println("Mutex Counter:", counter)
}

// using sync.RWMutex for read-only access
func BenchmarkRWMutexRead(b *testing.B) {
	var counter int
	var mutex sync.RWMutex

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			mutex.RLock()
			value := counter
			mutex.RUnlock()
			_ = value
		}
	})
	fmt.Println("RWMutex Read Counter:", counter)
}

// using sync.RWMutex for write access
func BenchmarkRWMutexWrite(b *testing.B) {
	var counter int
	var mutex sync.RWMutex

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			mutex.Lock()
			counter++
			mutex.Unlock()
		}
	})
	fmt.Println("RWMutex Write Counter:", counter)
}
func main() {
	testing.Benchmark(BenchmarkMutex)
	testing.Benchmark(BenchmarkRWMutexRead)
	testing.Benchmark(BenchmarkRWMutexWrite)
}
