package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sync"
)

func processFile(filePath string, wg *sync.WaitGroup) {
	defer wg.Done()

	// Attempt to open the file
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening file %s: %v\n", filePath, err)
		return
	}
	defer file.Close() // Ensure the file is closed

	// Attempt to process the file
	_, err = ioutil.ReadAll(file) // Simulate processing by reading the file
	if err != nil {
		fmt.Printf("Error processing file %s: %v\n", filePath, err)
		return
	}

	fmt.Println("File processed successfully:", filePath)
}

func main() {
	var wg sync.WaitGroup

	// Process multiple files concurrently
	filePaths := []string{"example1.txt", "example2.txt", "example3.txt", "example4.txt"}
	for _, filePath := range filePaths {
		wg.Add(1)
		go processFile(filePath, &wg)
	}

	// Wait for all goroutines to complete before exiting
	wg.Wait()
	fmt.Println("All files processed.")
}
