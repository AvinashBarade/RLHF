package benchmarks

import (
	"sync"
	"testing"
)

// BenchmarkMutex measures the overhead of a sync.Mutex.
func BenchmarkMutex(b *testing.B) {
	var counter int
	var mutex sync.Mutex

	for i := 0; i < b.N; i++ {
		mutex.Lock()
		counter++
		mutex.Unlock()
	}
}

// BenchmarkRWMutexWrite measures the overhead of a sync.RWMutex for writing.
func BenchmarkRWMutexWrite(b *testing.B) {
	var counter int
	var rwMutex sync.RWMutex

	for i := 0; i < b.N; i++ {
		rwMutex.Lock()
		counter++
		rwMutex.Unlock()
	}
}

// BenchmarkRWMutexRead measures the overhead of a sync.RWMutex for reading.
func BenchmarkRWMutexRead(b *testing.B) {
	var counter int
	var rwMutex sync.RWMutex

	for i := 0; i < b.N; i++ {
		rwMutex.RLock()
		_ = counter // Simulate read access.
		rwMutex.RUnlock()
	}
}

// BenchmarkWaitGroup measures the overhead of a sync.WaitGroup.
func BenchmarkWaitGroup(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			defer wg.Done()
		}()
		wg.Wait()
	}
}
