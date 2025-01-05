package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

type ChatApplication struct {
	listeners  map[string]func(string)
	monitoring *MonitoringSystem
}

type MonitoringSystem struct {
	logger *log.Logger
}

func (ms *MonitoringSystem) LogCallbackExecution(callbackName string, startTime time.Time, err error) {
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	msg := fmt.Sprintf("Callback %s executed in %s", callbackName, duration)
	if err != nil {
		msg += fmt.Sprintf(" with error: %v", err)
	}
	ms.logger.Println(msg)
}

func NewMonitoringSystem(logger *log.Logger) *MonitoringSystem {
	return &MonitoringSystem{logger}
}

func (c *ChatApplication) RegisterListener(listenerName string, callback func(string)) {
	c.listeners[listenerName] = callback
}

func (c *ChatApplication) SendMessage(message string) {
	for listenerName, callback := range c.listeners {
		go func(listenerName string, callback func(string)) {
			startTime := time.Now()
			defer c.monitoring.LogCallbackExecution(listenerName, startTime, nil)
			callback(message)
		}(listenerName, callback)
	}
}

func NewChatApplication(monitoring *MonitoringSystem) *ChatApplication {
	return &ChatApplication{
		listeners:  make(map[string]func(string)),
		monitoring: monitoring,
	}
}

func main() {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)
	monitoring := NewMonitoringSystem(logger)
	chatApp := NewChatApplication(monitoring)

	// Register callbacks
	chatApp.RegisterListener("User1", func(message string) {
		fmt.Println("User1 received message:", message)
		time.Sleep(1 * time.Second) // Simulate processing time
	})

	chatApp.RegisterListener("User2", func(message string) {
		fmt.Println("User2 received message:", message)
		// Simulate an error
		panic("Something went wrong!")
	})

	// Send a message to all listeners
	chatApp.SendMessage("Hello, this is a test message!")
}
