package pages

import (
	"log"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Cyan903/ws-gui/gui"
)

func Request() gui.Page {
	request := gui.CreatePage("Request")
	infoText := widget.NewLabel("Not connected")
	inputText := widget.NewMultiLineEntry()
	reqUrl := widget.NewEntry()

	inputText.SetPlaceHolder("Some request info...")
	reqUrl.SetPlaceHolder("wss://example.com")

	connectBtn := widget.NewButton("Connect", func() {
		infoText.SetText("Awaiting connection...")
	})

	inpForm := &widget.Form{
		Items: []*widget.FormItem{{
			Widget: container.NewAdaptiveGrid(2,
				connectBtn, container.NewCenter(infoText),
			),
		}, {
			Widget: reqUrl,
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
