package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// SignalHandler defines an interface for signal handlers.
type SignalHandler interface {
	HandleSignal(sig os.Signal)
}

// GracefulShutdownHandler implements SignalHandler for graceful shutdown.
type GracefulShutdownHandler struct {
	// Add any additional fields required for graceful shutdown here.
}

// HandleSignal handles OS signals.
func (gsh *GracefulShutdownHandler) HandleSignal(sig os.Signal) {
	switch sig {
	case syscall.SIGINT, syscall.SIGTERM:
		fmt.Println("Received signal:", sig)
		fmt.Println("Gracefully shutting down...")
		// Perform graceful shutdown operations here.
		os.Exit(0)
	default:
		fmt.Println("Unrecognized signal:", sig)
	}
}

// RegisterSignalHandlers registers signal handlers.
func RegisterSignalHandlers(handlers ...SignalHandler) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		for sig := range sigChan {
			for _, handler := range handlers {
				handler.HandleSignal(sig)
			}
		}
	}()
}

func main() {
	// Create a GracefulShutdownHandler.
	gsh := &GracefulShutdownHandler{}

	// Register the GracefulShutdownHandler with RegisterSignalHandlers.
	RegisterSignalHandlers(gsh)

	// The rest of your application code goes here.
	fmt.Println("Application started...")
	select {}
}
