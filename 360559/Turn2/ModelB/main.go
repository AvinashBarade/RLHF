package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/go-playground/validator/v10"
)

type Address struct {
	Street string `json:"street" validate:"required"`
	City   string `json:"city" validate:"required"`
}

type Person struct {
	Name    string   `json:"name" validate:"required"`
	Age     int      `json:"age" validate:"required,gte=0"`
	Address *Address `json:"address" validate:"required"`
}

func main() {
	// Create a validator instance
	validate := validator.New()

	// Sample nested JSON data
	data := []byte(`{"name": "Alice", "age":25, "address": {"city": "New York"}}`)

	var person Person
	err := json.Unmarshal(data, &person)
	if err != nil {
		log.Fatalf("Error unmarshaling JSON: %v", err)
	}

	// Validate the unmarshalled struct
	err = validate.Struct(person)
	if err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			// Print all the validation errors
			for _, ve := range ve {
				fmt.Printf("Error: Field: %v, Tag: %v, Value: %v\n", ve.Field(), ve.Tag(), ve.Value())
			}
		} else {
			// Handle any other errors
			log.Fatalf("Validation error: %v", err)
		}
	} else {
		fmt.Printf("JSON data is valid: %+v\n", person)
	}
}
