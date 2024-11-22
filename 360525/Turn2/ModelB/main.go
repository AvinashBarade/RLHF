package main

import (
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

func Double(x int) int {
	return x * 2
}

func Capitalize(s string) string {
	return strings.Title(s)
}

func HandleForm(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		numbers := r.Form["numbers"]
		var transformedNumbers []int
		for _, numStr := range numbers {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				http.Error(w, "Invalid number input", http.StatusBadRequest)
				return
			}
			transformedNumbers = append(transformedNumbers, Double(num))
		}

		data := struct {
			TransformedNumbers []int
		}{
			TransformedNumbers: transformedNumbers,
		}

		templates := template.Must(template.ParseFiles("template.html"))
		templates.Funcs(template.FuncMap{"double": Double, "capitalize": Capitalize})
		if err := templates.Execute(w, data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		http.ServeFile(w, r, "template.html")
	}
}

func main() {
	http.HandleFunc("/", HandleForm)
	http.ListenAndServe(":8080", nil)
}
