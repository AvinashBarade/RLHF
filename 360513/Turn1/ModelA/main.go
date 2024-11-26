package main

import (
	"net/http"

	"github.com/gorilla/websocket"
)

// CustomHandshaker implements a custom websocket handshake that checks for a special header.
type CustomHandshaker struct {
	websocket.Upgrader
	requiredHeader string
}

// Handshake performs the websocket handshake and checks for the required header.
func (h *CustomHandshaker) Handshake(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	// Check if the required header is present
	if r.Header.Get(h.requiredHeader) == "" {
		http.Error(w, "Missing required header", http.StatusUnauthorized)
		return nil, nil
	}

	// Perform the standard websocket handshake
	return h.Upgrader.Upgrade(w, r, nil)
}

func main() {
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		// Create a new CustomHandshaker with the required header
		handshaker := CustomHandshaker{
			Upgrader:       websocket.Upgrader{},
			requiredHeader: "X-Custom-Token",
		}

		// Perform the custom handshake
		conn, err := handshaker.Handshake(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		defer conn.Close()

		// Handle the websocket connection
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}
				return
			}

			// Echo the message back to the client
			err = conn.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
	})

	http.ListenAndServe(":8080", nil)
}
