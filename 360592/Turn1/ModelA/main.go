package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

func main() {
	// Simulating a list of tasks
	tasks := []string{"Task 1", "Task 2", "Task 3"}

	for _, task := range tasks {
		// Send a task to the consumer service
		url := "http://localhost:8081/process-task"
		data := fmt.Sprintf("task=%s", task)
		resp, err := http.Post(url, "application/x-www-form-urlencoded", strings.NewReader(data))
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
