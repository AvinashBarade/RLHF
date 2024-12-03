package main

import (
	"fmt"
	"time"
)

func produceData(data chan<- int) {
	for i := 0; i < 10; i++ {
		data <- i
		time.Sleep(100 * time.Millisecond)
	}
	close(data)
}

func consumeData(data <-chan int) {
	for value := range data {
		fmt.Println("Consumed:", value)
	}
}

func main() {
	data := make(chan int, 3) // Buffer size of 3

	go produceData(data)
	go consumeData(data)

	<-data // Wait for the data channel to be closed
}
