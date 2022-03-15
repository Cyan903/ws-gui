package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

type Page struct {
	Container *fyne.Container
	Title     string
}

func CreatePage(title string) Page {
	p := Page{
		container.NewCenter(),
		title,
	}

	return p
}

func (p Page) AddItem(add fyne.CanvasObject) {
	p.Container.Add(add)
}
