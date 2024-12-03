package main

import (
	"fmt"
	"net/http"
)

func handleRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handleRequest)
	fmt.Println("Server starting on http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
