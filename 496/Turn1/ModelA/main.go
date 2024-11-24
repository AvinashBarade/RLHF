package main

import (
	"fmt"
	"runtime"
	"sync"
)

// Pool for reusing Request objects
type Request struct {
	data string
}

var requestPool *sync.Pool

func init() {
	requestPool = &sync.Pool{
		New: func() interface{} {
			return &Request{}
		},
	}
}

func handleRequest() {
	// Get a Request object from the pool
	req := requestPool.Get().(*Request)
	defer requestPool.Put(req)

	// Simulate processing the request
	req.data = "Processed data"
	fmt.Println(req.data)
}

func main() {
	// Start GC debug logs
	runtime.SetGCDebug(1)

	for i := 0; i < 10000; i++ {
		handleRequest()
	}

	// Forcibly run a GC cycle to see debug logs
	runtime.GC()
}
