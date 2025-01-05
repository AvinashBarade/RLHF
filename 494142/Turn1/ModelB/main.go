package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
)

// scientificData represents the scientific data
type scientificData map[string]int

func (s *scientificData) Update(key string, value int) {
	(*s)[key] = value
}

func (s *scientificData) Read(key string) int {
	return (*s)[key]
}

// scanner simulates the user interaction by receiving updates
func scanner(updates chan string, wg *sync.WaitGroup) {
	for {
		var key, value string
		fmt.Print("Enter key and value separated by space: ")
		fmt.Scanln(&key, &value)
		updates <- fmt.Sprintf("%s=%s", key, value)
		if key == "quit" {
			close(updates)
			wg.Done()
			return
		}
	}
}

func main() {
	var wg sync.WaitGroup
	data := scientificData{"sensor1": 0, "sensor2": 0, "sensor3": 0}
	updates := make(chan string)
	wg.Add(1)

	//monitor the user interaction and update the data
	go scanner(updates, &wg)

	//update the data
	go func() {
		for update := range updates {
			keyValue := update[strings.Index(update, "=")+1:]
			key := keyValue[:strings.Index(keyValue, "=")]
			value, _ := strconv.Atoi(keyValue[strings.Index(keyValue, "=")+1:])
			data.Update(key, value)
			fmt.Println("\nData Updated:", data)
		}
	}()

	wg.Wait()
	fmt.Println("\nProgram Ended!")
}
