package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type UserProfile struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Bio   string `json:"bio"` // New field
}

func v2GetUserProfileHandler(w http.ResponseWriter, r *http.Request) {
	userProfile := UserProfile{
		Name:  "John Doe",
		Email: "john@example.com",
		Bio:   "Software Engineer with a passion for Go.",
	}
	json.NewEncoder(w).Encode(userProfile)
}

func v1GetUserProfileHandler(w http.ResponseWriter, r *http.Request) {
	// Compatibly return only the old fields
	userProfile := UserProfile{
		Name:  "John Doe",
		Email: "john@example.com",
		Bio:   "", // Default to empty string to maintain backward compatibility
	}
	json.NewEncoder(w).Encode(userProfile)
}

func main() {
	r := mux.NewRouter()

	// Version 1 API
	r.HandleFunc("/v1/profile", v1GetUserProfileHandler).Methods("GET")

	// Version 2 API
	r.HandleFunc("/v2/profile", v2GetUserProfileHandler).Methods("GET")

	fmt.Println("Server listening on :8080")
	http.ListenAndServe(":8080", r)
}
