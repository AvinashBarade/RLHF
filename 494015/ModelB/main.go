package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

// Session struct to hold session information
type Session struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
}

var (
	sessions          = make(map[string]Session) // Map to hold session ID to session data
	sessionsMut       = sync.Mutex{}             // Mutex to ensure thread safety
	sessionExpiryTime = 10 * time.Second         // Session expiry time in seconds
)

func deleteExpiredSessions() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for range ticker.C {
		sessionsMut.Lock()
		for sessionID, session := range sessions {
			if time.Since(session.CreatedAt) > sessionExpiryTime {
				delete(sessions, sessionID)
				fmt.Printf("Session %s expired and removed.\n", sessionID)
			}
		}
		sessionsMut.Unlock()
	}
}

// Function to add a session
func addSession(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	if username == "" {
		http.Error(w, "Username is required", http.StatusBadRequest)
		return
	}
	sessionsMut.Lock()
	defer sessionsMut.Unlock()
	sessionID := fmt.Sprintf("session-%d", time.Now().UnixNano())
	newSession := Session{
		ID:        sessionID,
		Username:  username,
		CreatedAt: time.Now(),
	}
	sessions[sessionID] = newSession
	bytes, err := json.MarshalIndent(newSession, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(bytes)
}

// Function to remove a session
func removeSession(w http.ResponseWriter, r *http.Request) {
	sessionID := r.URL.Query().Get("sessionID")
	if sessionID == "" {
		http.Error(w, "Session ID is required", http.StatusBadRequest)
		return
	}
	sessionsMut.Lock()
	defer sessionsMut.Unlock()
	if _, ok := sessions[sessionID]; !ok {
		http.Error(w, "Session not found", http.StatusNotFound)
		return
	}
	delete(sessions, sessionID)
	fmt.Fprintln(w, "Session removed")
}

// Function to monitor sessions and compute growth rate
func monitorSessions() {
	previousCount := 0
	ticker := time.NewTicker(time.Second * 5)
	defer ticker.Stop()
	for range ticker.C {
		sessionsMut.Lock()
		currentCount := len(sessions)
		sessionsMut.Unlock()
		if currentCount > 0 {
			growthRate := float64(currentCount-previousCount) / float64(currentCount) * 100
			fmt.Printf("Active sessions: %d, Growth rate: %.2f%%\n", currentCount, growthRate)
		}
		previousCount = currentCount
	}
}
func main() {
	go deleteExpiredSessions() // Start a goroutine to delete expired sessions
	go monitorSessions()       // Start a goroutine to monitor sessions

	http.HandleFunc("/add-session", addSession)
	http.HandleFunc("/remove-session", removeSession)
	fmt.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server error:", err)
	}
}
