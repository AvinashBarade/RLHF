package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Task struct {
	Name string `json:"task"`
}

func main() {
	// Simulating a list of tasks
	tasks := []string{"Task 1", "Task 2", "Task 3"}

	for _, task := range tasks {
		// Send a task to the consumer service
		taskJSON, err := json.Marshal(Task{Name: task})
		if err != nil {
			fmt.Println("Error marshaling task:", err)
			continue
		}

		url := "http://localhost:8081/process-task"
		resp, err := http.Post(url, "application/json", bytes.NewReader(taskJSON))
		if err != nil {
			fmt.Println("Error sending task:", err)
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			fmt.Println("Failed to process task:", resp.Status)
			continue
		}

		fmt.Println("Task sent successfully:", task)
		time.Sleep(1 * time.Second) // Simulate a delay
	}
}
