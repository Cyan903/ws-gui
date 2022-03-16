package pages

import (
	"fyne.io/fyne/v2/widget"
	"github.com/Cyan903/ws-gui/gui"
)

func Response() gui.Page {
	response := gui.CreatePage("Response")
	resText := widget.NewMultiLineEntry()

	response.AddItem(resText)

	return response
}
