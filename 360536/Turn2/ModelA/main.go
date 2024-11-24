package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"

	"github.com/golang/heap"
)

// Task struct represents a task with a priority and a work function.
type Task struct {
	ID       string
	Work     func()
	Priority int // Higher priority tasks are processed first
}

// TaskQueue implements the heap interface for task prioritization.
type TaskQueue []Task

// Implement the Less method for the heap interface.
func (tq TaskQueue) Less(i, j int) bool {
	return tq[i].Priority > tq[j].Priority
}

// Implement the Swap method for the heap interface.
func (tq TaskQueue) Swap(i, j int) {
	tq[i], tq[j] = tq[j], tq[i]
}

// Implement the Len method for the heap interface.
func (tq TaskQueue) Len() int {
	return len(tq)
}

// Implement the Push method for the heap interface.
func (tq *TaskQueue) Push(x interface{}) {
	*tq = append(*tq, x.(Task))
}

// Implement the Pop method for the heap interface.
func (tq *TaskQueue) Pop() interface{} {
	old := *tq
	n := len(old)
	old[n-1], old[0] = old[0], old[n-1]
	x := old[n-1]
	*tq = old[:n-1]
	return x
}

// TaskRunner struct to handle task execution with prioritization.
type TaskRunner struct {
	taskQueue      TaskQueue
	numWorkers     int
	working        sync.WaitGroup
	exitChannel    chan bool
	completedTasks int
}

// NewTaskRunner initializes a new TaskRunner with a given number of workers.
func NewTaskRunner(numWorkers int) *TaskRunner {
	return &TaskRunner{
		taskQueue:      TaskQueue{},
		numWorkers:     numWorkers,
		working:        sync.WaitGroup{},
		exitChannel:    make(chan bool),
		completedTasks: 0,
	}
}

// Start begins task processing with the specified number of workers.
func (tr *TaskRunner) Start() {
	for i := 0; i < tr.numWorkers; i++ {
		go tr.worker()
	}
}

// Stop stops task processing and waits for all tasks to complete.
func (tr *TaskRunner) Stop() {
	tr.exitChannel <- true
	tr.working.Wait()
	fmt.Printf("Completed tasks: %d\n", tr.completedTasks)
}

// AddTask adds a new task to the task queue.
func (tr *TaskRunner) AddTask(task Task) {
	heap.Push(&tr.taskQueue, task)
}

// worker is a goroutine that processes tasks from the task queue.
func (tr *TaskRunner) worker() {
	defer tr.working.Done()
	for {
		select {
		case <-tr.exitChannel:
			// Exit signal received, stop processing tasks
			return
		default:
			if !heap.Empty(&tr.taskQueue) {
				tr.working.Add(1)
				task := heap.Pop(&tr.taskQueue).(Task)

				// Process the task
				task.Work()
				tr.completedTasks++
				fmt.Printf("Worker %d completed task: %s with priority: %d\n", rand.Intn(10)+1, task.ID, task.Priority)
				tr.working.Done()
			}
		}
		time.Sleep(time.Millisecond) // Limit CPU usage
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	tr := NewTaskRunner(10) // Run with 10 worker goroutines
	tr.Start()

	// Add a set of tasks to the queue with random priorities
	for i := 0; i < 100; i++ {
		taskID := strconv.Itoa(i)
		tr.AddTask(Task{
			ID: taskID,
			Work: func() {
				// Simulate some work
				time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
			},
			Priority: rand.Intn(100), // Random priority between 0 and 99
		})
	}

	// Gracefully stop the task runner
	time.Sleep(3 * time.Second) // Allow tasks to complete
	tr.Stop()
}
