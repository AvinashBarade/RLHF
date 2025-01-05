package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

type Client struct {
	conn     net.Conn
	addr     string
	reader   *bufio.Reader
	writer   *bufio.Writer
	username string
}

var clients map[string]*Client
var callbacks map[string]func(*Client, []byte)

func main() {
	clients = make(map[string]*Client)
	callbacks = make(map[string]func(*Client, []byte))

	registerCallback("message", handleMessage)
	registerCallback("private_message", handlePrivateMessage)

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
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	client := &Client{
		conn:     conn,
		addr:     conn.RemoteAddr().String(),
		reader:   reader,
		writer:   writer,
		username: "Guest",
	}

	defer func() {
		conn.Close()
		delete(clients, client.username)
	}()

	fmt.Fprintf(client.writer, "Enter your username: ")
	client.writer.Flush()

	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				log.Printf("Client %s disconnected", client.addr)
				return
			}
			log.Fatalf("Client %s read error: %v", client.addr, err)
			return
		}

		username := strings.TrimSpace(string(line))
		if username != "" && clients[username] == nil {
			client.username = username
			clients[username] = client
			log.Printf("Client %s connected with username %s", client.addr, client.username)
			break
		}

		fmt.Fprintf(client.writer, "Username already taken or invalid. Please try again: ")
		client.writer.Flush()
	}

	handleUser(client)
}

func handleUser(client *Client) {
	for {
		line, _, err := client.reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				log.Printf("Client %s disconnected", client.addr)
				return
			}
			log.Fatalf("Client %s read error: %v", client.addr, err)
			return
		}

		parts := strings.Split(string(line), " ")
		switch len(parts) {
		case 2:
			handlePrivateMessage(client, []byte(fmt.Sprintf("%s %s", client.username, parts[1])))
		default:
			handleMessage(client, line)
		}
	}
}

func handleMessage(client *Client, message []byte) {
	log.Printf("Handling message from %s: %s", client.username, message)
	for _, c := range clients {
		if c != client {
			c.writer.Write(append([]byte(client.username+": "), message...))
			c.writer.Flush()
		}
	}
}

func handlePrivateMessage(client *Client, message []byte) {
	log.Printf("Handling private message from %s: %s", client.username, message)
	parts := strings.SplitN(string(message), " ", 2)
	if len(parts) != 2 {
		fmt.Fprintf(client.writer, "Invalid private message format. Use: private <recipient> <message>\n")
		client.writer.Flush()
		return
	}

	recipient := strings.TrimSpace(parts[0])
	if clients[recipient] == nil {
		fmt.Fprintf(client.writer, "Recipient not found: %s\n", recipient)
		client.writer.Flush()
		return
	}

	clients[recipient].writer.Write(append([]byte(client.username+": "), []byte(parts[1])...))
	clients[recipient].writer.Flush()
}

func registerCallback(name string, callback func(*Client, []byte)) {
	callbacks[name] = callback
}
