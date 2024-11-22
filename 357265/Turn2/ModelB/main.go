package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func fetchAPI() string {
	resp, err := http.Get("http://localhost:8080/api/endpoint") // Replace with your API URL
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}
	return string(body)
}

func main() {
	a := app.New()
	w := a.NewWindow("API Explorer")

	apiButton := widget.NewButton("Fetch API", func() {
		response := fetchAPI()
		if response != "" {
			dialog.ShowInformation("API Response", response, w)
		} else {
			dialog.ShowError(fmt.Errorf("Error fetching API"), w)
		}
	})

	w.SetContent(container.New(layout.NewCenter(), apiButton))
	w.ShowAndRun()
}
