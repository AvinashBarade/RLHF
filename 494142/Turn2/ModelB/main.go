package main

import (
	"fmt"
	"regexp"
	"strconv"
	"sync"
	"time"
)

var (
	temperatureData = make(map[string]float64)
	mutex           = sync.RWMutex{}
	dataUpdates     = make(chan string)
)

func isValidSensorID(sensorID string) bool {
	// You can customize the regex pattern to match your specific requirements for sensor ID validation
	pattern := `^[A-Za-z0-9_]{3,20}$`
	return regexp.MustCompile(pattern).MatchString(sensorID)
}

func isValidTemperature(temp string) bool {
	value, err := strconv.ParseFloat(temp, 64)
	return err == nil && value >= -273.15 && value <= 1000.0
}

func updateData() {
	for {
		sensorID := input("Enter sensor ID: ")
		temp := input("Enter temperature: ")

		// Validate sensor ID
		if !isValidSensorID(sensorID) {
			fmt.Println("Invalid sensor ID. Please enter a valid alphanumeric ID between 3 to 20 characters long.")
			continue
		}

		// Validate temperature
		if !isValidTemperature(temp) {
			fmt.Println("Invalid temperature. Please enter a valid number between -273.15 and 1000.0.")
			continue
		}

		tempValue, _ := strconv.ParseFloat(temp, 64)
		mutex.Lock()
		temperatureData[sensorID] = tempValue
		mutex.Unlock()
		dataUpdates <- sensorID
	}
}

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
		case <-time.After(2 * time.Second):
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

func main() {
	go updateData()
	go displayData()
	select {}
}
