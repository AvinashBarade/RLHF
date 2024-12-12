package main

import (
	"container/list"
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// Constants for API rate limiting
const (
	rateLimit  = 5  // Requests per second
	burstLimit = 10 // Requests that can burst
	cacheSize  = 20 // LRU Cache Size
)

var (
	bucket       = make(chan time.Time, burstLimit)
	cacheMetrics = Metrics{}
	cache        = NewLRUCache(cacheSize)
)

type Metrics struct {
	cacheHits   int
	cacheMisses int
	rejections  int
	mu          sync.Mutex
}

// LRUCache implementation
type LRUCache struct {
	capacity int
	cache    map[string]*list.Element
	order    *list.List
	mutex    sync.Mutex
}

type entry struct {
	key   string
	value string
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		cache:    make(map[string]*list.Element),
		order:    list.New(),
	}
}

func (c *LRUCache) Get(key string) (string, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if elem, found := c.cache[key]; found {
		c.order.MoveToFront(elem) // move accessed item to the front
		return elem.Value.(*entry).value, true
	}

	return "", false
}

func (c *LRUCache) Put(key, value string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if elem, found := c.cache[key]; found {
		c.order.MoveToFront(elem)
		elem.Value.(*entry).value = value
		return
	}

	elem := c.order.PushFront(&entry{key: key, value: value})
	c.cache[key] = elem

	if c.order.Len() > c.capacity {
		c.removeOldest()
	}
}

func (c *LRUCache) removeOldest() {
	if elem := c.order.Back(); elem != nil {
		c.order.Remove(elem)
		entry := elem.Value.(*entry)
		delete(c.cache, entry.key)
	}
}

func (m *Metrics) IncrementCacheHit() {
	m.mu.Lock()
	m.cacheHits++
	m.mu.Unlock()
}

func (m *Metrics) IncrementCacheMiss() {
	m.mu.Lock()
	m.cacheMisses++
	m.mu.Unlock()
}

func (m *Metrics) IncrementRejection() {
	m.mu.Lock()
	m.rejections++
	m.mu.Unlock()
}

func main() {
	go fillBucket()

	ctx, cancel := context.WithCancel(context.Background())

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-quit
		cancel()
	}()

	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("key-%d", i)
		go fetchData(ctx, key)
	}

	select {
	case <-time.After(5 * time.Second):
		printMetrics()
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
		return // Exit if context is cancelled
	case <-bucket:
		// Rate limited, proceed with the request
	default:
		cacheMetrics.IncrementRejection()
		log.Printf("Rate limit exceeded for key: %s\n", key)
		return
	}

	// Simulate caching
	if value, ok := cache.Get(key); ok {
		cacheMetrics.IncrementCacheHit()
		log.Printf("Data fetched from cache for key: %s, Value: %s\n", key, value)
		return
	}
	cacheMetrics.IncrementCacheMiss()

	// Simulate fetching data from an external source
	sleepDuration := time.Duration(rand.Intn(500)) * time.Millisecond
	time.Sleep(sleepDuration)
	value := fmt.Sprintf("Fetched data for key: %s after %v", key, sleepDuration)

	// Simulate caching the fetched data
	cache.Put(key, value)

	log.Printf("Data fetched and cached for key: %s, Value: %s\n", key, value)
}

func printMetrics() {
	fmt.Printf("Cache Hits: %d, Cache Misses: %d, Rate Limit Rejections: %d\n",
		cacheMetrics.cacheHits, cacheMetrics.cacheMisses, cacheMetrics.rejections)
}
