package main

import (
	"html/template"
	"log"
	"net/http"
)

// PageData holds the data for our page
type PageData struct {
	Title   string
	Numbers []int
}

// templateFuncs is a FuncMap that contains our custom functions
var templateFuncs = template.FuncMap{
	"double": func(i int) int {
		return i * 2
	},
}

func main() {
	// Create a new template and parse the template file
	tmpl, err := template.New("index.html").Funcs(templateFuncs).ParseFiles("index.html")
	if err != nil {
		log.Fatalln(err)
	}

	// Define our page data
	pageData := PageData{
		Title:   "Hello, World!",
		Numbers: []int{1, 2, 3, 4, 5},
	}

	// Handle the HTTP request
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Execute the template and write the result to the HTTP response
		if err := tmpl.Execute(w, pageData); err != nil {
			log.Println(err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	})

	// Start the HTTP server
	log.Println("Server starting on port :8080")
	log.Fatalln(http.ListenAndServe(":8080", nil))
}
