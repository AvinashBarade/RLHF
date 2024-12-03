package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
)

func searchFile(fileName string, pattern string, results chan<- string) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	if strings.Contains(string(data), pattern) {
		results <- fileName
	}
}

func main() {
	var wg sync.WaitGroup
	results := make(chan string)

	pattern := "go"
	searchDir := "/path/to/your/search/directory"

	// Calculate the number of goroutines based on the number of CPU cores
	numGoroutines := runtime.NumCPU()

	filepath.Walk(searchDir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		wg.Add(1)
		go searchFile(path, pattern, results)
		return nil
	})

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Println("Found file:", result)
	}
}
