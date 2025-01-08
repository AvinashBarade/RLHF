package main

import (
	"fmt"
	"log"
	"os"
	"reflect"
)

// FilterPositiveNumbers filters out negative numbers from a slice of integers, handles mixed types, and logs skipped elements.
func FilterPositiveNumbers(numbers interface{}) ([]int, error) {
	value := reflect.ValueOf(numbers)

	if value.Kind() != reflect.Slice {
		log.Printf("ERROR: Input type is %v. Expected slice.", value.Kind())
		return nil, fmt.Errorf("input type is %v. Expected slice.", value.Kind())
	}

	// Filter positive integers while handling mixed types
	filtered := []int{}
	skippedElements := []struct {
		Index int
		Value interface{}
	}{}
	for i := 0; i < value.Len(); i++ {
		element := value.Index(i).Interface()
		switch v := element.(type) {
		case int:
			if v >= 0 {
				filtered = append(filtered, v)
			}
		case float64:
			if int64(v) >= 0 {
				filtered = append(filtered, int(v))
			}
		case int64:
			if v >= 0 {
				filtered = append(filtered, int(v))
			}
		case uint, uint32, uint64:
			if int64(v) >= 0 {
				filtered = append(filtered, int(v))
			}
		default:
			skippedElements = append(skippedElements, struct {
				Index int
				Value interface{}
			}{i, element})
			log.Printf("WARNING: Skipping non-integer value at index %d: %v", i, element)
		}
	}

	return filtered, nil
}

func setupLogger() {
	logFile, err := os.OpenFile("application.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Failed to set up log file: %v\n", err)
		os.Exit(1)
	}
	log.SetOutput(logFile)
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.SetPrefix("[FilterLogger] ")
}

func main() {
	setupLogger()

	// Example input slice with mixed types
	input := []interface{}{10, "hello", -5, 3.14, 0, 23, -7, 8, "world", 13}

	log.Println("Starting filtering process")

	filtered, err := FilterPositiveNumbers(input)
	if err != nil {
		log.Printf("ERROR: %v\n", err)
		return
	}

	log.Printf("Filtering complete. Input: %v, Filtered: %v\n", input, filtered)
	fmt.Println("Filtered slice:", filtered)
}
