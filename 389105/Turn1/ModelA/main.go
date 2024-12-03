package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const numTasks = 20
const maxTaskDuration = 2 * time.Second

// Task represents a unit of work to be processed
type Task struct {
	Id int
}

func generateTask() Task {
	return Task{Id: rand.Intn(numTasks) + 1}
}

func worker(id int, workerQueue chan Task, globalTasks chan Task, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case task := <-workerQueue:
			fmt.Printf("Worker %d is processing task %d\n", id, task.Id)
			time.Sleep(time.Duration(rand.Intn(int(maxTaskDuration))) * time.Millisecond)

		default:
			// Steal a task from the global queue if local queue is empty
			select {
			case task := <-globalTasks:
				fmt.Printf("Worker %d stole task %d from global queue\n", id, task.Id)
				workerQueue <- task
			default:
				// No tasks in the global queue, check for completion
				wg.Wait()
				return
			}
		}
	}
}

func main() {
	numWorkers := 5
	workerQueues := make([]chan Task, numWorkers)
	globalTasks := make(chan Task)
	var wg sync.WaitGroup

	rand.Seed(time.Now().UnixNano())

	wg.Add(numTasks)
	for i := 0; i < numWorkers; i++ {
		workerQueues[i] = make(chan Task, 100)
		go worker(i+1, workerQueues[i], globalTasks, &wg)
	}

	// Generate and distribute tasks to workers dynamically
	for i := 0; i < numTasks; i++ {
		task := generateTask()
		// Distribute tasks randomly among workers
		workerQueues[rand.Intn(numWorkers)] <- task
	}

	// Close the global tasks queue
	close(globalTasks)

	wg.Wait()
	fmt.Println("All tasks completed.")
}
