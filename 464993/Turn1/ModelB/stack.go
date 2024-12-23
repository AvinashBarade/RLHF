package main

import (
	"fmt"
	"os"
	"strconv"
)

// Stack is a custom data type that wraps a slice of integers
type Stack struct {
	items []int
}

// NewStack creates a new empty Stack
func NewStack() *Stack {
	return &Stack{}
}

// IsEmpty checks if the stack is empty
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

// Push adds an element to the top of the stack
func (s *Stack) Push(element int) {
	s.items = append(s.items, element)
}

// Pop removes and returns the top element from the stack
func (s *Stack) Pop() int {
	if s.IsEmpty() {
		panic("Stack is empty")
	}
	topElement := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return topElement
}

// Peek returns the top element of the stack without removing it
func (s *Stack) Peek() int {
	if s.IsEmpty() {
		panic("Stack is empty")
	}
	return s.items[len(s.items)-1]
}

func main() {
	stack := NewStack()

	// Command-line tool usage: go run stack.go push 10 push 20 pop peek push 30 peek
	args := os.Args[1:]

	for _, arg := range args {
		switch arg {
		case "push":
			if len(args) < 2 {
				fmt.Println("Invalid command. Usage: push <number>")
				continue
			}
			num, err := strconv.Atoi(args[1])
			if err != nil {
				fmt.Println("Invalid number:", args[1])
				continue
			}
			stack.Push(num)
			args = args[1:]
		case "pop":
			if stack.IsEmpty() {
				fmt.Println("Stack is empty")
			} else {
				fmt.Println("Popped element:", stack.Pop())
			}
		case "peek":
			if stack.IsEmpty() {
				fmt.Println("Stack is empty")
			} else {
				fmt.Println("Top element:", stack.Peek())
			}
		default:
			fmt.Println("Invalid command:", arg)
		}
	}
}
