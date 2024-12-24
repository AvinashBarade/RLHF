package main

import (
	"fmt"
	"sync"
)

type Event struct {
	name string
}

type Observer interface {
	Notify(event *Event)
}

type Subject struct {
	observers []Observer
	mutex     sync.Mutex
}

func (s *Subject) Attach(observer Observer) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.observers = append(s.observers, observer)
}

func (s *Subject) Detach(observer Observer) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	for i, o := range s.observers {
		if o == observer {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			return
		}
	}
}

func (s *Subject) NotifyObservers(event *Event) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	for _, observer := range s.observers {
		observer.Notify(event)
	}
}

type Worker struct {
	id       int
	subject  *Subject
	doneChan chan struct{}
}

func NewWorker(id int, subject *Subject) *Worker {
	return &Worker{
		id:       id,
		subject:  subject,
		doneChan: make(chan struct{}),
	}
}

func (w *Worker) Start() {
	go func() {
		defer close(w.doneChan)
		for {
			select {
			case <-w.doneChan:
				fmt.Printf("Worker %d: Stopped\n", w.id)
				return
			default:
				fmt.Printf("Worker %d: Working...\n", w.id)
				w.subject.NotifyObservers(&Event{name: "WorkerStarted"})
			}
		}
	}()
}

func (w *Worker) Stop() {
	close(w.doneChan)
}

type EventHandler struct {
	id int
}

func (h *EventHandler) Notify(event *Event) {
	fmt.Printf("Handler %d: Received event: %s\n", h.id, event.name)
}

func main() {
	subject := &Subject{}

	worker1 := NewWorker(1, subject)
	worker2 := NewWorker(2, subject)

	handler1 := &EventHandler{id: 1}
	handler2 := &EventHandler{id: 2}

	subject.Attach(handler1)
	subject.Attach(handler2)

	worker1.Start()
	worker2.Start()

	// Simulate some work and then stop the workers
	fmt.Println("Doing some work...")

	worker1.Stop()
	worker2.Stop()

	// Wait for workers to finish
	<-worker1.doneChan
	<-worker2.doneChan

	subject.Detach(handler1)
	subject.Detach(handler2)
}
