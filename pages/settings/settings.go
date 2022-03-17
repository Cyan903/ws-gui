package settings

import (
	"log"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Cyan903/ws-gui/gui"
)

func Settings() gui.Page {
	settings := gui.CreatePage("Settings")

	checkUpdates := widget.NewButton("Check for updates", func() {})
	clearHistory := widget.NewButton("Clear connection history", func() {})

	clearOnExit := widget.NewCheck("Clear on exit", func(n bool) {
		log.Println(n)
	})

	settings.AddItem(container.NewGridWithColumns(2,
		clearOnExit, widget.NewLabel(""), checkUpdates, clearHistory,
	))

	return settings
}
