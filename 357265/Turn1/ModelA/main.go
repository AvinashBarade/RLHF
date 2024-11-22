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
		dialog.ShowError("Error", err.Error(), fyne.NewApp())
		return ""
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		dialog.ShowError("Error", err.Error(), fyne.NewApp())
		return ""
	}

	return string(body)
}

func main() {
	app := fyne.NewApp()
	win := app.NewWindow("API Explorer")

	apiButton := widget.NewButton("Fetch API", func() {
		response := fetchAPI()
		if response != "" {
			dialog.ShowInformation("API Response", response, win)
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

	win.SetContent(content)
	win.ShowAndRun()
}
