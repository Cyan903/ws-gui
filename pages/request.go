package pages

import (
	"log"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Cyan903/ws-gui/gui"
)

func Request() gui.Page {
	request := gui.CreatePage("Request")
	infoText := widget.NewLabel("Connected!")
	inputText := widget.NewMultiLineEntry()

	inputText.SetText("input text test")

	connectBtn := widget.NewButton("Disconnect", func() {
		log.Println("Disconnected...")
	})

	inpForm := &widget.Form{
		Items: []*widget.FormItem{{
			Widget: container.NewAdaptiveGrid(2,
				connectBtn, container.NewCenter(infoText),
			),
		}, {
			Widget: inputText,
		}},
		OnSubmit: func() {
			log.Println("Form submitted:", inputText.Text)
			log.Println("multiline:", inputText.Text)
		},
	}

	inpForm.SubmitText = "Send"
	request.AddItem(inpForm)

	return request
}
