package settings

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Cyan903/ws-gui/gui"
)

func Settings() gui.Page {
	settings := gui.CreatePage("Settings")

	checkUpdates := widget.NewButton("Check for updates", checkUpdates)
	clearHistory := widget.NewButton("Clear connection history", clearConnectionHistory)
	clearConsole := widget.NewButton("Clear console", clearConsole)
	clearOnDisconnect := widget.NewCheck("Clear on connect/disconnect", clearDisconnectFn)

	settings.AddItem(container.NewGridWithColumns(2,
		clearOnDisconnect, clearConsole, checkUpdates, clearHistory,
	))

	return settings
}
