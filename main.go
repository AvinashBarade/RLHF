package main

import (
	"fmt"
	"sync"
	"time"
)

// Task represents a job to be done by a Goroutine, with a priority level.
type Task struct {
	id       int
	priority int // 1 for high, 0 for low
}

// Worker function consumes tasks from priority queues.
func worker(id int, highPriorityTasks <-chan Task, lowPriorityTasks <-chan Task, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case task, ok := <-highPriorityTasks:
			if !ok {
				break // Channel closed, exit worker
			}
			processTask(task, id, "high")
		case task, ok := <-lowPriorityTasks:
			if !ok {
				break // Channel closed, exit worker
			}
			processTask(task, id, "low")
		}
	}
}

func processTask(task Task, workerId int, priority string) {
	fmt.Printf("Worker %d processing %s priority task %d\n", workerId, priority, task.id)
	time.Sleep(time.Second) // Simulate work with a sleep
}

func main() {
	const numWorkers = 3
	const numHighPriorityTasks = 5
	const numLowPriorityTasks = 5

	// Channels for tasks
	highPriorityTasks := make(chan Task)
	lowPriorityTasks := make(chan Task)

	var wg sync.WaitGroup

	// Start worker Goroutines
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i, highPriorityTasks, lowPriorityTasks, &wg)
	}

	// Add high-priority tasks
	for i := 0; i < numHighPriorityTasks; i++ {
		highPriorityTasks <- Task{id: i, priority: 1}
	}

	// Add low-priority tasks
	for i := 0; i < numLowPriorityTasks; i++ {
		lowPriorityTasks <- Task{id: i + numHighPriorityTasks, priority: 0}
	}

	// Close task channels after sending all tasks
	close(highPriorityTasks)
	close(lowPriorityTasks)

	// Wait for all workers to complete
	wg.Wait()
	fmt.Println("All tasks completed.")
}
