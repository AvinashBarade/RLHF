package main

import (
	"bufio"
	"fmt"
	"os"
)

func processLine(line string, results chan<- int) {
	count := len(line)
	results <- count
}

func main() {
	file, err := os.Open("example.log")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	results := make(chan int)

	for scanner.Scan() {
		go processLine(scanner.Text(), results)
	}

	go func() {
		var total int
		for count := range results {
			total += count
		}
		fmt.Println("Total length of lines:", total)
		close(results)
	}()

	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file:", err)
	}
}
