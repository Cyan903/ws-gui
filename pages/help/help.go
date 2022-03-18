package help

import (
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"github.com/Cyan903/ws-gui/gui"
	"github.com/skratchdot/open-golang/open"
)

func fallback() *fyne.Container {
	c := container.NewCenter(container.NewAdaptiveGrid(1,
		widget.NewLabel("Could not find help markdown! Check static folder?"),
		container.NewAdaptiveGrid(2,
			widget.NewButton("View in browser", func() {
				open.Run(helpURL)
			}),

			widget.NewButton("Download", func() {
				go downloadHelp()
			}),
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
		sli := strings.Split(md, "\n")
		items := container.NewAdaptiveGrid(1)

		for i := range sli {
			items.Add(widget.NewRichTextFromMarkdown(sli[i]))
		}

		help.AddItem(widget.NewRichTextFromMarkdown(md))
	}

	return help
}
