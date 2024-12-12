package main

import (
	"fmt"
	"log"
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

// ProfileGoroutines profiles the running Goroutines
func ProfileGoroutines(profilerChan chan Job) {
	for job := range profilerChan {
		// Gather Goroutine info
		buf := make([]byte, 1<<20)    // Create a buffer
		n := runtime.Stack(buf, true) // Get stack trace
		fmt.Printf("Profiling Goroutine %d - Stack Trace:\n%s\n", job.ID, buf[:n])

		// Could also collect or process any additional info about the job
	}
}

func main() {
	profilerChan := make(chan Job)
	go ProfileGoroutines(profilerChan) // Goroutine to monitor other Goroutines

	// Start CPU profiling
	f, err := os.Create("cpu.prof")
	if err != nil {
		log.Fatal("could not create CPU profile:", err)
	}
	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatal("could not start CPU profile:", err)
	}
	defer pprof.StopCPUProfile()

	// Simulate a number of jobs being processed
	for i := 0; i < 5; i++ {
		job := Job{ID: i, Duration: 2 * time.Second} // Create job
		go func(job Job) {
			profilerChan <- job // Send job to profiler
			ProcessJob(job)     // Process the job
		}(job)
	}

	time.Sleep(5 * time.Second) // Allow time for jobs to spawn and for profiling to occur
	close(profilerChan)         // Cleanly close the channel

	// Start memory profiling
	runtime.GC() // Ensure all garbage is collected
	f, err := os.Create("mem.prof")
	if err != nil {
		log.Fatal("could not create memory profile:", err)
	}
	if err := pprof.WriteHeapProfile(f); err != nil {
		log.Fatal("could not write memory profile:", err)
	}
	f.Close()
}
