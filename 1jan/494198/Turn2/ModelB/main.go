package main

import (
	"fmt"
	"log"
	"net"
	"strings"
)

type Client struct {
	conn net.Conn
	addr string
	name string
}

var clients map[string]*Client = make(map[string]*Client)
var callbacks map[string]func(*Client, []byte) = make(map[string]func(*Client, []byte))

func main() {
	registerCallback("message", handleMessage)
	registerCallback("private", handlePrivateMessage)
	ln, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("Error listening: %v", err)
	}
	defer ln.Close()
	log.Println("Chat server started on :9090")
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("Error accepting connection:", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	client := &Client{conn: conn, addr: conn.RemoteAddr().String()}
	clients[client.addr] = client
	defer func() {
		delete(clients, client.addr)
		conn.Close()
	}()
	log.Println("New client connected:", client.addr)
	for {
		message, err := readMessage(client)
		if err != nil {
			log.Println("Error reading message:", err)
			return
		}
		handleMessage(client, []byte(message))
	}
}

func readMessage(client *Client) (string, error) {
	buf := make([]byte, 1024)
	n, err := client.conn.Read(buf)
	if err != nil {
		return "", err
	}
	return string(buf[:n]), nil
}

func handleMessage(client *Client, message []byte) {
	msg := string(message)
	if strings.HasPrefix(msg, "/private ") {
		handlePrivateMessage(client, message)
	} else {
		broadcastMessage(client, message)
	}
}

func broadcastMessage(client *Client, message []byte) {
	msg := fmt.Sprintf("%s: %s", client.name, string(message))
	for _, c := range clients {
		c.conn.Write([]byte(msg + "\n"))
	}
}

func handlePrivateMessage(client *Client, message []byte) {
	msg := string(message)
	parts := strings.SplitN(msg, " ", 3)
	if len(parts) != 3 {
		client.conn.Write([]byte("Invalid private message format. Use /private <recipient> <message>.\n"))
		return
	}
	recipient := parts[1]
	privateMessage := parts[2]
	c, ok := clients[recipient]
	if !ok {
		client.conn.Write([]byte("Recipient not found.\n"))
		return
	}
	privateMsg := fmt.Sprintf("(private from %s) %s", client.name, privateMessage)
	c.conn.Write([]byte(privateMsg + "\n"))
}

func registerCallback(name string, callback func(*Client, []byte)) {
	callbacks[name] = callback
}
