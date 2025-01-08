package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

func countWordsInFile(ctx context.Context, filename string, wordCounts chan<- map[string]int) {
	file, err := os.Open(filename)
	if err != nil {
		log.Printf("Error opening file %s: %v\n", filename, err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	counts := make(map[string]int)

	for scanner.Scan() {
		select {
		case <-ctx.Done():
			log.Println("Word count task cancelled for", filename)
			return
		default:
			words := strings.Fields(scanner.Text())
			for _, word := range words {
				counts[word]++
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Error reading file %s: %v\n", filename, err)
	}

	wordCounts <- counts
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	filenames := []string{"example1.txt", "example2.txt", "example3.txt"}
	wordCounts := make(chan map[string]int, len(filenames))

	var wg sync.WaitGroup
	wg.Add(len(filenames))

	for _, filename := range filenames {
		go func(filename string) {
			defer wg.Done()
			countWordsInFile(ctx, filename, wordCounts)
		}(filename)
	}

	go func() {
		wg.Wait()
		close(wordCounts)
	}()

	totalCounts := make(map[string]int)
	for counts := range wordCounts {
		for word, count := range counts {
			totalCounts[word] += count
		}
	}

	fmt.Println("Final word counts:")
	for word, count := range totalCounts {
		fmt.Printf("%s: %d\n", word, count)
	}
}
