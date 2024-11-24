package main

import (
	"fmt"
	"math/rand"
	"sort"
	"sync"
	"time"

	"github.com/dgryski/go-metro"
)

const (
	numReplicas = 16 // Number of replicas per shard
)

type HashRing struct {
	mu       sync.RWMutex
	replicas []string
	keys     []uint64
}

func (r *HashRing) AddReplica(key string) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.replicas = append(r.replicas, key)
	r.updateKeys()
}

func (r *HashRing) RemoveReplica(key string) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for i, replica := range r.replicas {
		if replica == key {
			r.replicas = append(r.replicas[:i], r.replicas[i+1:]...)
			break
		}
	}
	r.updateKeys()
}

func (r *HashRing) FindReplica(key string) string {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if len(r.keys) == 0 {
		return "" // Ring is empty
	}

	hashValue := metro.Hash64String(key)
	index := sort.SearchUint64s(r.keys, hashValue)
	if index == len(r.keys) {
		index = 0
	}
	return r.replicas[index/numReplicas]
}

func (r *HashRing) updateKeys() {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.keys = make([]uint64, 0, len(r.replicas)*numReplicas)
	for _, key := range r.replicas {
		for i := 0; i < numReplicas; i++ {
			replicaKey := fmt.Sprintf("%s:%d", key, i)
			hashValue := metro.Hash64String(replicaKey)
			r.keys = append(r.keys, hashValue)
		}
	}
	sort.Slice(r.keys, func(i, j int) bool { return r.keys[i] < r.keys[j] })
}

func main() {
	ring := &HashRing{}
	ring.AddReplica("shard1")
	ring.AddReplica("shard2")
	ring.AddReplica("shard3")
	ring.AddReplica("shard4")

	keys := []string{"user1", "user2", "user3", "user4", "user5", "product100", "product1000", "order99999"}
	rand.Seed(time.Now().UnixNano())

	fmt.Println("Mapping keys to shards:")
	for _, key := range keys {
		shard := ring.FindReplica(key)
		fmt.Printf("%s -> %s\n", key, shard)
	}

	// Add and remove shards to demonstrate dynamism
	ring.AddReplica("shard5")
	fmt.Println("\nShard 5 added. Updated mappings:")
	for _, key := range keys {
		shard := ring.FindReplica(key)
		fmt.Printf("%s -> %s\n", key, shard)
	}

	ring.RemoveReplica("shard3")
	fmt.Println("\nShard 3 removed. Updated mappings:")
	for _, key := range keys {
		shard := ring.FindReplica(key)
		fmt.Printf("%s -> %s\n", key, shard)
	}
}
