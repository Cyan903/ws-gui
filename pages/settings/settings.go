package settings

import (
	"net/url"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Cyan903/ws-gui/gui"
)

func Settings() gui.Page {
	settings := gui.CreatePage("Settings")

	clearHistory := widget.NewButton("Clear connection history", clearConnectionHistory)
	clearConsole := widget.NewButton("Clear console", clearConsole)
	clearOnDisconnect := widget.NewCheck("Clear on connect/disconnect", clearDisconnectFn)
	homeUrl, _ := url.Parse("https://github.com/Cyan903/ws-gui/")

	settings.AddItem(container.NewGridWithColumns(2,
		container.NewCenter(clearOnDisconnect), container.NewCenter(widget.NewHyperlink("Source code", homeUrl)), 
		clearConsole, clearHistory,
	))

	return settings
}
