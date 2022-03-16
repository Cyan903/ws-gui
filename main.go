package main

import (
	"fyne.io/fyne/v2/app"
	"github.com/Cyan903/ws-gui/gui"
	"github.com/Cyan903/ws-gui/pages"
)

func main() {
	app := app.New()
	win := app.NewWindow("ws-gui Go")

	tabs := gui.CreateTabs([]gui.Page{
		pages.Response(),
		pages.Request(),
		pages.History(),
		pages.Settings(),
	})

	win.SetContent(tabs)
	win.ShowAndRun()
}
