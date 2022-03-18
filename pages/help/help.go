package help

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Cyan903/ws-gui/gui"
)

func fallback() *fyne.Container {
	c := container.NewCenter(container.NewAdaptiveGrid(1, 
		widget.NewLabel("Could not find help markdown! Check static folder?"),
		container.NewAdaptiveGrid(2,
			widget.NewButton("View in browser", func() {}),
			widget.NewButton("Download", func() {}),
		),
	))

	return c
}

func Help() gui.Page {
	help := gui.CreatePage("Help")
	md := readHelp()

	if md == "" {
		help.AddItem(fallback())
	} else {
		help.AddItem(widget.NewRichTextFromMarkdown(md))
	}

	return help
}
