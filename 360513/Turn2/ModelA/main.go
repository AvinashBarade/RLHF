package main

import (
	"log"
	"net/http"
	"strings"

	"golang.org/x/net/websocket"
)

type customHandshake struct {
	userID string
}

func (ch *customHandshake) Handshake(config *websocket.Config, req *http.Request) error {
	// Extract the X-User-ID header from the request
	userIDs := req.Header.Get("X-User-ID")

	// Check if the header is missing or if it's empty
	if strings.TrimSpace(userIDs) == "" {
		log.Println("Error: X-User-ID header is missing or empty")
		http.Error(req.ResponseWriter, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return websocket.ErrProtocolViolation
	}

	// Extract the first user ID from a comma-separated list
	userIDs = strings.Split(userIDs, ",")[0]
	log.Printf("User ID extracted from header: %s\n", userIDs)
	ch.userID = userIDs
	return nil
}

func (ch *customHandshake) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	websocket.Handler(func(ws *websocket.Conn) {
		log.Println("New WebSocket connection from:", ch.userID)
		for {
			var msg string
			err := websocket.Message.Receive(ws, &msg)
			if err != nil {
				if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
					log.Println("Client disconnected normally")
				} else {
					log.Println("Error receiving message:", err)
				}
				break
			}
			log.Println("Received message:", msg)
			err = websocket.Message.Send(ws, "Echo: "+msg)
			if err != nil {
				log.Println("Error sending message:", err)
				break
			}
		}
		ws.Close()
		log.Println("WebSocket connection closed")
	}).ServeHTTP(w, r)
}

func main() {
	handshake := &customHandshake{}
	log.Println("Server started on :8080")
	err := http.ListenAndServe(":8080", handshake)
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}
