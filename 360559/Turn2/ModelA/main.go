package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/go-playground/validator/v10"
)

type Address struct {
	Street  string `json:"street" validate:"required"`
	City    string `json:"city" validate:"required"`
	ZipCode string `json:"zip_code" validate:"required"`
}

type Person struct {
	Name    string   `json:"name" validate:"required"`
	Age     int      `json:"age" validate:"required,gt=0"`
	Address *Address `json:"address" validate:"required"`
}

func main() {
	// Sample JSON data
	jsonData := []byte(`
	{
		"name": "John Doe",
		"age": 30,
		"address": {
			"street": "123 Main St",
			"city": "Anytown",
			"zip_code": "12345"
		}
	}
	`)

	// Deserialize JSON to a Person struct
	var person Person
	err := json.Unmarshal(jsonData, &person)
	if err != nil {
		log.Fatalf("Error unmarshaling JSON: %v", err)
	}

	// Validate the Person struct
	validate := validator.New()
	err = validate.Struct(person)
	if err != nil {
		fmt.Println("Validation errors:")
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Printf("%s: %s\n", err.Field(), err.Tag())
		}
		return
	}

	// If no errors, print the validated person information
	fmt.Printf("Validated Person: %+v\n", person)
}
