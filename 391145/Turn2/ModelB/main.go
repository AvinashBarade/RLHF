package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

// Constants for rate limiting
const (
	rateLimit  = 5  // Requests per second
	burstLimit = 10 // Requests that can burst
)

var (
	cache   *cache.Cache
	bucket  = make(chan time.Time, burstLimit)
	metrics = struct {
		cacheHits   int
		cacheMisses int
		rateLimits  int
		sync.Mutex
	}{}
	wg sync.WaitGroup
)

func init() {
	cache = cache.New(5*time.Minute, 30*time.Minute)
	rand.Seed(time.Now().UnixNano())
}

func main() {
	go fillBucket()

	ctx, cancel := context.WithCancel(context.Background())

	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("key-%d", i)
		wg.Add(1)
		go func() {
			defer wg.Done()
			fetchData(ctx, key)
		}()
	}

	// Graceful shutdown
	go func() {
		<-time.After(5 * time.Second) // Simulate some running time
		log.Println("Initiating graceful shutdown...")
		cancel()
		wg.Wait()
		log.Println("All goroutines stopped. Exiting...")
	}()

	select {
	case <-ctx.Done():
	}
}

func fillBucket() {
	for {
		time.Sleep(time.Second / time.Duration(rateLimit))
		bucket <- time.Now()
	}
}

func fetchData(ctx context.Context, key string) {
	select {
	case <-ctx.Done():
		return // Exit if the context is cancelled
	case <-bucket:
		// Rate limited, proceed with the request
	default:
		metrics.Lock()
		metrics.rateLimits++
		metrics.Unlock()
		log.Printf("Rate limit exceeded for key: %s\n", key)
		return
	}

	// Simulate caching with LRU
	if item, ok := cache.Get(key); ok {
		cache.Touch(key) // Mark as recently used
		metrics.Lock()
		metrics.cacheHits++
		metrics.Unlock()
		log.Printf("Data fetched from cache for key: %s, Value: %v\n", key, item)
		return
	}

	// Simulate fetching data from an external source
	sleepDuration := time.Duration(rand.Intn(500)) * time.Millisecond
	time.Sleep(sleepDuration)
	value := fmt.Sprintf("Fetched data for key: %s after %v", key, sleepDuration)

	// Simulate caching the fetched data
	cache.Set(key, value, cache.DefaultExpiration)
	metrics.Lock()
	metrics.cacheMisses++
	metrics.Unlock()

	log.Printf("Data fetched and cached for key: %s, Value: %s\n", key, value)
}

func printMetrics() {
	metrics.Lock()
	defer metrics.Unlock()
	fmt.Printf("Metrics: Cache Hits=%d, Cache Misses=%d, Rate Limit Rejections=%d\n",
		metrics.cacheHits, metrics.cacheMisses, metrics.rateLimits)
}
