package main

import (
	"io"
	"log"
	"net"
	"time"
)

type Client struct {
	conn net.Conn
	addr string
}

var clients []*Client
var callbacks map[string]func(*Client, []byte)

func main() {
	callbacks = make(map[string]func(*Client, []byte))

	// Register a message handling callback
	registerCallback("message", handleMessage)

	// Set up a TCP listener
	ln, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("Could not listen on :9090: %v", err)
	}
	defer ln.Close()

	log.Printf("Listening on :9090...")

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatalf("Accept failed: %v", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	client := &Client{
		conn: conn,
		addr: conn.RemoteAddr().String(),
	}
	clients = append(clients, client)
	defer func() {
		clients = removeClient(clients, client)
		conn.Close()
	}()

	log.Printf("Client %s connected", client.addr)

	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				log.Printf("Client %s disconnected", client.addr)
				return
			}
			log.Fatalf("Client %s read error: %v", client.addr, err)
			return
		}

		message := string(buf[:n])
		log.Printf("Received message from %s: %s", client.addr, message)

		for _, client := range clients {
			// Check if the client is not the sender
			if client.addr != client.addr {
				client.conn.Write([]byte(message))
			}
		}

		// Check if there's a callback registered for this message
		if callback, ok := callbacks["message"]; ok {
			callback(client, buf[:n])
		}
	}
}

func removeClient(clients []*Client, client *Client) []*Client {
	for i, c := range clients {
		if c == client {
			return append(clients[:i], clients[i+1:]...)
		}
	}
	return clients
}

func registerCallback(name string, callback func(*Client, []byte)) {
	callbacks[name] = callback
}

func handleMessage(client *Client, message []byte) {
	log.Printf("Handling message for client %s", client.addr)
	startTime := time.Now()
	log.Printf("Processing message...")
	time.Sleep(100 * time.Millisecond) // Simulate processing time
	endTime := time.Now()
	log.Printf("Message processed in %v", endTime.Sub(startTime))
}
