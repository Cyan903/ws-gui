package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func CreateTabs(pages []Page) fyne.CanvasObject {
	tabs := container.NewAppTabs()

	for _, p := range pages {
		tabs.Append(container.NewTabItem(p.Title, p.Container))
	}

	return tabs
}
