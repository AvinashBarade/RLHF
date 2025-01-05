package main

import (
	"fmt"
	"regexp"
	"sync"
	"time"
)

var (
	temperatureData = make(map[string]float64) // A map to hold sensor data
	mutex           = sync.RWMutex{}           // Mutex for thread-safe operations on the map
	dataUpdates     = make(chan string)        // Channel to trigger data updates
)

// updateData simulates user interaction to update sensor data with input validation
func updateData() {
	for {
		sensorID := input("Enter sensor ID (non-empty string): ")
		if isValidSensorID(sensorID) {
			temp, err := inputFloat("Enter temperature (positive number): ")
			if err == nil && temp > 0 {
				// Update the temperature data in a thread-safe manner
				mutex.Lock()
				temperatureData[sensorID] = temp
				mutex.Unlock()

				// Signal that the data has been updated
				dataUpdates <- sensorID
			} else if err != nil {
				fmt.Println("Invalid temperature, please enter a valid positive number.")
			} else {
				fmt.Println("Temperature must be a positive number.")
			}
		} else {
			fmt.Println("Invalid sensor ID, please enter a non-empty string.")
		}
	}
}

// displayData continuously displays the current state of the dataset
func displayData() {
	for sensorID := range temperatureData {
		fmt.Printf("Sensor %s: %.2f°C\n", sensorID, temperatureData[sensorID])
	}

	for {
		select {
		case sensorID := <-dataUpdates:
			mutex.RLock()
			fmt.Printf("Sensor %s updated to %.2f°C\n", sensorID, temperatureData[sensorID])
			mutex.RUnlock()
		case <-time.After(2 * time.Second): // Refresh data every 2 seconds
			mutex.RLock()
			for sensorID, temp := range temperatureData {
				fmt.Printf("Sensor %s: %.2f°C\n", sensorID, temp)
			}
			mutex.RUnlock()
		}
	}
}

func input(prompt string) string {
	fmt.Print(prompt)
	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		panic(err)
	}
	return input
}

func inputFloat(prompt string) (float64, error) {
	fmt.Print(prompt)
	var num float64
	_, err := fmt.Scanf("%f", &num)
	return num, err
}

// isValidSensorID validates that the sensor ID is a non-empty string
func isValidSensorID(id string) bool {
	return regexp.MustCompile(`^[a-zA-Z0-9_\-]+$`).MatchString(id) && id != ""
}

func main() {
	go updateData()  // Start goroutine for user interactions
	go displayData() // Start goroutine to display data updates

	select {} // Blocks indefinitely to keep the program running
}
