package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

const APIVersion = "v1"

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/"+APIVersion+"/users", getUsersHandler).Methods("GET")
	r.HandleFunc("/api/"+APIVersion+"/users/{id}", getUserHandler).Methods("GET")
	r.HandleFunc("/api/"+APIVersion+"/users", createUserHandler).Methods("POST")
	http.ListenAndServe(":8080", r)
}

func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World! You are using API version %s", APIVersion)
}

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World! You are using API version %s", APIVersion)
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World! You are using API version %s", APIVersion)
}
