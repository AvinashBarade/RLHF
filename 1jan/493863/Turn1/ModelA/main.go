package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("example.txt")
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close() // Ensure file is closed even if there's an error

	data, err := os.ReadAll(file)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	fmt.Println(string(data))
}
