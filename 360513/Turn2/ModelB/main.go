package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"golang.org/x/net/websocket"
)

type customHandshake struct {
	userID string
}

func (ch *customHandshake) Handshake(config *websocket.Config, req *http.Request) error {
	// Perform custom validation of the request
	userID := req.Header.Get("X-User-ID")
	if userID == "" {
		// Log error on server side
		log.Println("Error: X-User-ID header is missing")
		return fmt.Errorf("User ID header is required")
	}

	// Add additional validation for user ID format if needed
	// For example, checking for valid UUID or email format

	ch.userID = userID
	return nil
}

func (ch *customHandshake) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	websocket.Handler(func(ws *websocket.Conn) {
		defer ws.Close()

		// Access the custom handshake data
		userID := ch.userID
		fmt.Println("New connection from:", userID)

		// Handle the WebSocket connection
		for {
			var msg string
			err := websocket.Message.Receive(ws, &msg)
			if err != nil {
				// Log error on server side
				log.Println("Error reading message:", err)
				break
			}

			fmt.Println("Received message:", msg)
			err = websocket.Message.Send(ws, "Echo: "+msg)
			if err != nil {
				// Log error on server side
				log.Println("Error sending message:", err)
				break
			}
		}
	}).ServeHTTP(w, r)
}

func main() {
	handshake := &customHandshake{}
	http.Handle("/ws", handshake)

	// Start the HTTP server with a graceful shutdown
	srv := &http.Server{
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Server started on :8080")
	err := srv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatal("Error starting server:", err)
	}
}
