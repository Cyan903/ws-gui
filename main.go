package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/Cyan903/ws-gui/gui"
	. "github.com/Cyan903/ws-gui/pages"
)

func main() {
	app := app.New()
	win := app.NewWindow("ws-gui Go")

	tabs := gui.CreateTabs([]gui.Page{
		Pages["request"](),
		Pages["response"](),
		Pages["history"](),
		Pages["settings"](),
	})

	win.SetContent(tabs)
	win.Resize(fyne.NewSize(600, 300))
	win.ShowAndRun()
}
