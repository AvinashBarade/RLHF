package main

import (
	"html/template"
	"net/http"
	"strings"
)

func Capitalize(s string) string {
	if s == "" {
		return ""
	}
	return strings.Title(s)
}

func main() {
	// Create a new template set
	templates := template.Must(template.ParseFiles("template.html"))

	// Define some data to be used by the template
	type Person struct {
		Name  string
		Email string
	}
	data := Person{"john doe", "john.doe@example.com"}

	// Register the custom function with the template
	templates.Funcs(template.FuncMap{"capitalize": Capitalize})

	// Serve the content
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := templates.Execute(w, data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	http.ListenAndServe(":8080", nil)
}
