package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type SignalHandleFunc func()
type SignalResult struct {
	Error error
	Data  interface{}
}

type SignalHandler struct {
	Name string
	Func SignalHandleFunc
}

var customSignals map[os.Signal]*SignalHandler
var results map[os.Signal]SignalResult
var wg sync.WaitGroup
var signalHandlerMutex sync.Mutex

func init() {
	customSignals = make(map[os.Signal]*SignalHandler)
	results = make(map[os.Signal]SignalResult)
}
func RegisterSignal(sig os.Signal, name string, handler SignalHandleFunc) {
	signalHandlerMutex.Lock()
	defer signalHandlerMutex.Unlock()

	// Check if signal is already registered
	_, ok := customSignals[sig]
	if ok {
		fmt.Printf("Signal '%s' (%v) is already registered.\n", name, sig)
		return
	}

	// Save the new signal handler
	customSignals[sig] = &SignalHandler{
		Name: name,
		Func: handler,
	}

	// Block the signal
	err := syscall.Sigprocmask(syscall.SIG_BLOCK, &syscall.Sigset_t{Set: []uint64{uint64(sig)}}, nil)
	if err != nil {
		fmt.Printf("Error blocking signal '%s': %v\n", name, err)
		return
	}

	// Unblock the signal in a goroutine to handle it asynchronously
	go handleSignal(sig)
}
func handleSignal(sig os.Signal) {
	for {
		select {
		case <-signal.Notify(make(chan os.Signal, 1), sig):
			handler, ok := customSignals[sig]
			if !ok {
				fmt.Printf("Received signal '%s' (%v), but no handler is registered.\n", sig, sig)
				return
			}

			wg.Add(1)
			defer wg.Done()
			handler.Func()
			results[sig] = SignalResult{Error: nil, Data: nil}
			fmt.Printf("Signal '%s' handled successfully.\n", handler.Name)
		}
	}
}
func GracefulShutdown() {
	fmt.Println("Graceful shutdown initiated...")
	// Perform any cleanup operations here
}

func main() {
	// Register a custom signal named "USR1"
	RegisterSignal(syscall.SIGUSR1, "USR1", func() {
		fmt.Println("Received SIGUSR1, data processing initiated...")
		// Do some custom work here
	})

	// Register a custom signal named "USR2"
	RegisterSignal(syscall.SIGUSR2, "USR2", func() {
		fmt.Println("Received SIGUSR2, custom operation triggered.")
		// Perform another custom task here
	})

	// Handle SIGTERM for graceful shutdown
	signal.Notify(make(chan os.Signal, 1), syscall.SIGTERM)
	go func() {
		<-signal.Notify(make(chan os.Signal, 1), syscall.SIGTERM)
		GracefulShutdown()
	}()

	// Wait for signals
	fmt.Println("Awaiting signals...")
	wg.Wait()

	// Process signal results
	fmt.Println("Signal Handling Complete:")
	for sig, result := range results {
		fmt.Printf("Signal: %v, Error: %v, Data: %v\n", sig, result.Error, result.Data)
	}
}
