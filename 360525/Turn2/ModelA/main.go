package main

import (
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

func Capitalize(s string) string {

	if s == "" {

		return ""

	}

	return strings.Title(s)

}

func Double(numbers []string) []string {

	doubledNumbers := make([]string, len(numbers))

	for i, number := range numbers {

		if num, err := strconv.Atoi(number); err == nil {

			doubledNumbers[i] = strconv.Itoa(num * 2)

		} else {

			doubledNumbers[i] = "Invalid input"

		}

	}
	return doubledNumbers

}

func main() {
	// Create a new template set
	templates := template.Must(template.ParseFiles("template.html"))

	// Define the data structure to be passed to the template
	type FormData struct {
		Numbers        []string
		DoubledNumbers []string
	}

	// Register the custom functions with the template
	templates.Funcs(template.FuncMap{
		"capitalize": Capitalize,
		"double":     Double,
	})

	// Serve the content
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Initialize form data
		data := FormData{}

		// Check if the form was submitted
		if r.Method == "POST" {
			// Get form values
			if numbers := r.FormValue("numbers"); numbers != "" {
				// Split the input string into a slice of strings
				data.Numbers = strings.Split(numbers, ",")
				// Transform the numbers using the double function
				data.DoubledNumbers = Double(data.Numbers)
			}
		}

		// Execute the template with the form data
		if err := templates.Execute(w, data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	http.ListenAndServe(":8081", nil)
}
