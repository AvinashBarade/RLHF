package main

import (
	"io/ioutil"
	"net/http"

	"fyne.io/fyne"
	"fyne.io/fyne/container"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

func fetchAPI() string {
	resp, err := http.Get("http://localhost:8080/api/endpoint") // Replace with your API URL
	if err != nil {
		dialog.ShowError("Error Fetching API", err.Error(), app) // Pass the app object
		return ""
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		dialog.ShowError("Error Reading Response", err.Error(), app)
		return ""
	}

	return string(body)
}

func main() {
	app := fyne.NewApp()

	apiButton := widget.NewButton("Fetch API", func() {
		response := fetchAPI()
		if response != "" {
			dialog.ShowInformation("API Response", response, app)
		}
	})

	layout := layout.NewVBox(layout.NewSpacer(), apiButton)
	content := container.NewBorder(
		"",
		"",
		"",
		layout,
		"",
	)

	win := app.NewWindow("API Explorer")
	win.SetContent(content)
	win.ShowAndRun()
}
