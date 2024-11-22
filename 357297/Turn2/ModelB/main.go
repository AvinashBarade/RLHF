package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Hello struct {
	Message string `json:"message"`
}

type HelloWithTimestamp struct {
	Hello
	Timestamp int64 `json:"timestamp"`
}

func v1HelloHandler(w http.ResponseWriter, r *http.Request) {
	response := Hello{Message: "Hello from v1 API"}
	json.NewEncoder(w).Encode(response)
}

func v2HelloHandler(w http.ResponseWriter, r *http.Request) {
	// New feature: Add a timestamp to the response
	timestamp := getCurrentTimestamp()
	response := HelloWithTimestamp{
		Hello:     Hello{Message: "Hello from v2 API"},
		Timestamp: timestamp,
	}
	json.NewEncoder(w).Encode(response)
}

func getCurrentTimestamp() int64 {
	// Replace this with actual timestamp generation logic
	return 1609459200 // January 1, 2021, 00:00:00 UTC
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/v1/hello", v1HelloHandler).Methods("GET")
	r.HandleFunc("/v2/hello", v2HelloHandler).Methods("GET")

	fmt.Println("Server listening on :8080")
	http.ListenAndServe(":8080", r)
}
