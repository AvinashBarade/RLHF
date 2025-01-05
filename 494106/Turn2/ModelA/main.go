package main

import (
	"fmt"
	"os"
)

// Define a nested data structure
type Employee struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Age     int     `json:"age"`
	Skills  []string `json:"skills"`
	Address map[string]string
}

func readFile(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file %s: %w", filePath, err)
	}
	defer file.Close()

	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", filePath, err)
	}

	return data, nil
}

func main() {
	// Create an instance of the Employee struct
	employee := Employee{
		ID:      1,
		Name:    "Alice Johnson",
		Age:     30,
		Skills:  []string{"Go", "JavaScript", "SQL"},
		Address: map[string]string{"Street": "123 Main St", "City": "New York", "Country": "USA"},
	}

	// Logging the Employee struct using %#v for detailed output
	fmt.Printf("Detailed Employee information: %#v\n", employee)

	// Logging the Employee struct using %+v for verbose output
	fmt.Printf("Verbose Employee information: %+v\n", employee)

	// Custom formatting using %s for specific fields
	formattedEmployee := fmt.Sprintf("Employee ID: %d, Name: %s, Age: %d, Skills: %s", employee.ID, employee.Name, employee.Age, employee.Skills)
	fmt.Printf("Custom formatted Employee information: %s\n", formattedEmployee)

	// Formatting a slice within the struct
	skillsString := fmt.Sprintf("Skills: %v\n", employee.Skills)
	fmt.Printf("Skills formatted: %s", skillsString)

	// Formatting a map within the struct
	addressString := fmt.Sprintf("Address: %v\n", employee.Address)
	fmt.Printf("Address formatted: %s", addressString)

	// Example of error handling
	filePath := "example.txt"
	data, err := readFile(filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("File content: %s\n", string(data))
}