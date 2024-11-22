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

func v1HelloHandler(w http.ResponseWriter, r *http.Request) {
	response := Hello{Message: "Hello from v1 API"}
	json.NewEncoder(w).Encode(response)
}

func v2HelloHandler(w http.ResponseWriter, r *http.Request) {
	response := Hello{Message: "Hello from v2 API"}
	json.NewEncoder(w).Encode(response)
}

func main() {
	r := mux.NewRouter()

	// Version 1 API
	r.HandleFunc("/v1/hello", v1HelloHandler).Methods("GET")

	// Version 2 API
	r.HandleFunc("/v2/hello", v2HelloHandler).Methods("GET")

	fmt.Println("Server listening on :8080")
	http.ListenAndServe(":8080", r)
}
