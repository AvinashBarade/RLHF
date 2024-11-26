package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Task struct {
	Name string `json:"name"`
}

func main() {
	tasks := []string{"Task 1", "Task 2", "Task 3"}
	for _, task := range tasks {
		t := Task{Name: task}
		err := sendTask(t)
		if err != nil {
			fmt.Println("Error sending task:", err)
			continue
		}
		fmt.Println("Task sent successfully:", task)
		time.Sleep(1 * time.Second) // Simulate delay
	}
}

func sendTask(t Task) error {
	taskJSON, err := json.Marshal(t)
	if err != nil {
		return err
	}
	resp, err := http.Post("http://localhost:8081/process-task", "application/json", bytes.NewBuffer(taskJSON))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var errResp struct {
			Error string `json:"error"`
		}
		err := json.NewDecoder(resp.Body).Decode(&errResp)
		if err != nil {
			return fmt.Errorf("failed to decode response: %w", err)
		}
		return fmt.Errorf("%s: %s", resp.Status, errResp.Error)
	}
	return nil
}
