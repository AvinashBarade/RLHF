package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Task struct {
	Name string `json:"name"`
}

func processTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, "Bad request: cannot decode task", http.StatusBadRequest)
		return
	}
	fmt.Println("Received task:", task.Name)
	// Simulate processing the task
	time.Sleep(2 * time.Second)
	fmt.Fprintln(w, "Task processed:", task.Name)
}

func main() {
	http.HandleFunc("/process-task", processTaskHandler)
	fmt.Println("Consumer service listening on http://localhost:8081")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
