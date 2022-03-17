package history

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/Cyan903/ws-gui/gui"
)

func History() gui.Page {
	history := gui.CreatePage("Connection History")

	data := [][]string{
		{"#1", "12:45am", "wss://", "example.com"},
		{"#1", "12:45am", "wss://", "example.com"},
		{"#1", "12:45am", "wss://", "example.com"},
		{"#1", "12:45am", "wss://", "example.com"},
	}

	list := widget.NewTable(
		func() (int, int) {
			return len(data), len(data[0])
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("example.com")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			if i.Col == 0 {
				o.(*widget.Label).SetText(fmt.Sprintf("#%d", i.Row+1))
				return
			}

			o.(*widget.Label).SetText(data[i.Row][i.Col])
		},
	)

	history.AddItem(list)

	return history
}
