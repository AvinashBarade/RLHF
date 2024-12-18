package main

import (
	"sync"
	"testing"
)

var wg sync.WaitGroup
var m sync.Mutex

func BenchmarkMutex(b *testing.B) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < b.N; i++ {
			m.Lock()
			m.Unlock()
		}
	}()
	wg.Wait()
}

func BenchmarkWaitGroup(b *testing.B) {
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
		}()
		wg.Wait()
	}
}

func BenchmarkNoSync(b *testing.B) {
	for i := 0; i < b.N; i++ {
	}
}

func TestMain(m *testing.M) {
	m.Run()
}
