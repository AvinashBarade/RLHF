package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

type customHandshake struct {
	userID string
}

func (ch *customHandshake) Handshake(config *websocket.Config, req *http.Request) error {
	// Perform custom validation of the request
	// For example, check for the presence of a specific cookie or header
	userID := req.Header.Get("X-User-ID")
	if userID == "" {
		return fmt.Errorf("User ID header is required")
	}
	ch.userID = userID
	return nil
}

func (ch *customHandshake) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Upgrade the HTTP connection to a WebSocket connection
	// and pass the customHandshake struct as the handshake config
	websocket.Handler(func(ws *websocket.Conn) {
		// Access the custom handshake data
		userID := ch.userID
		fmt.Println("New connection from:", userID)
		// Handle the WebSocket connection
		// For example, read messages from the client and send responses
		for {
			var msg string
			err := websocket.Message.Receive(ws, &msg)
			if err != nil {
				log.Println("Error reading message:", err)
				break
			}
			fmt.Println("Received message:", msg)
			err = websocket.Message.Send(ws, "Echo: "+msg)
			if err != nil {
				log.Println("Error sending message:", err)
				break
			}
		}
	}).ServeHTTP(w, r)
}

func main() {
	// Create a new custom handshake struct
	handshake := &customHandshake{}
	// Start the HTTP server
	http.Handle("/ws", handshake)
	log.Println("Server started on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}
