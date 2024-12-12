package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/pprof"
	"time"
)

// Job represents a task to be processed
type Job struct {
	ID       int
	Duration time.Duration
}

// ProcessJob simulates processing a job
func ProcessJob(job Job) {
	time.Sleep(job.Duration) // Simulate workload
}

// ProfileGoroutines profiles the running Goroutines and records memory stats
func ProfileGoroutines(profilerChan chan Job) {
	for job := range profilerChan {
		// Goroutine stack trace
		buf := make([]byte, 1<<20)    // Create a buffer
		n := runtime.Stack(buf, true) // Get stack trace
		fmt.Printf("Profiling Goroutine %d - Stack Trace:\n%s\n", job.ID, buf[:n])

		// Capture memory statistics
		var memStats runtime.MemStats
		runtime.ReadMemStats(&memStats)
		fmt.Printf("Memory Usage:\n Alloc = %v MiB\n TotalAlloc = %v MiB\n Sys = %v MiB\n NumGC = %v\n",
			memStats.Alloc/1024/1024,
			memStats.TotalAlloc/1024/1024,
			memStats.Sys/1024/1024,
			memStats.NumGC)

		// Optionally, you can also start CPU profiling here
		f, err := os.Create(fmt.Sprintf("cpu_profile_%d.prof", job.ID))
		if err != nil {
			fmt.Println("Could not create CPU profile:", err)
			return
		}
		if err := pprof.StartCPUProfile(f); err != nil {
			fmt.Println("Could not start CPU profile:", err)
			return
		}
		defer pprof.StopCPUProfile()

		// Process the job
		ProcessJob(job)
	}
}

func main() {
	profilerChan := make(chan Job)
	go ProfileGoroutines(profilerChan) // Goroutine to monitor other Goroutines

	// Simulate a number of jobs being processed
	for i := 0; i < 5; i++ {
		job := Job{ID: i, Duration: 2 * time.Second} // Create job
		go func(job Job) {
			profilerChan <- job // Send job to profiler
		}(job)
	}

	time.Sleep(10 * time.Second) // Allow time for jobs to spawn and for profiling to occur
	close(profilerChan)          // Cleanly close the channel
}
