// gui.go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/goki/gi/gi"
	"github.com/goki/gi/gimain"
	"github.com/goki/gi/giv"
)

func main() {
	gimain.Main(func() {
		gui := NewGUI()
		gui.Start()
	})
}

// GUI represents the main GUI application
type GUI struct {
	win  *gi.Window
	view *giv.TableView
}

// NewGUI creates a new GUI instance
func NewGUI() *GUI {
	g := &GUI{}
	return g
}

// Start initializes and shows the GUI
func (g *GUI) Start() {
	width := 800
	height := 600

	g.win = gi.NewMainWindow("User Explorer", "User Explorer", width, height)
	g.win.SetCloseClean(true, true)

	vp := g.win.WinViewport2D()
	up := vp.UpdateStart()

	mfr := gi.AddNewFrame(vp, "main-frame", gi.LayoutVert)
	mfr.SetPaddings(0.2, 0.2, 0.2, 0.2)

	btn := gi.AddNewButton(mfr, "get-users-btn")
	btn.SetText("Get Users")
	btn.ConnectClicked(func(recv, send ki.Ki, sig int64, data any) {
		g.getUsers()
	})

	g.view = giv.AddNewTableView(mfr, "users-table")
	g.view.SetStretchMaxWidth()
	g.view.SetStretchMaxHeight()
	g.view.SetColumns([]string{"ID", "Name"})

	vp.UpdateEndNoSig(up)

	g.win.Show()
}

func (g *GUI) getUsers() {
	// Call the GoLang REST API to get users
	resp, err := http.Get("http://localhost:8080/users")
	if err != nil {
		fmt.Println("Error getting users:", err)
		return
	}
	defer resp.Body.Close()

	// Decode the JSON response into a slice of User structs
	var users []User
	if err := json.NewDecoder(resp.Body).Decode(&users); err != nil {
		fmt.Println("Error decoding JSON response:", err)
		return
	}

	// Update the table view with the fetched users
	g.view.ClearRows()
	for _, user := range users {
		row := g.view.AddRow()
		row.SetValues(user.ID, user.Name)
	}
	g.view.Refresh()
}
