package request

import (
	"log"
	"net/url"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Cyan903/ws-gui/client"
	"github.com/Cyan903/ws-gui/gui"
)

func connectionButton(t *widget.Label, r *widget.Entry) *widget.Button {
	connectBtn := conButton{}
	client := client.Client{
		Url: url.URL{},

		OnMessage: func(msg []byte, err error) {
			if err != nil {
				return
			}

			log.Println(string(msg))
		},

		OnError: func(err error) {
			connectBtn.Disconnect(err)
		},
	}

	connectBtn = conButton{
		connected: false,
		lbl:       t,
		con:       client,
		btn:       widget.NewButton("Connect", func() {}),
	}

	connectBtn.btn.OnTapped = toggleConnect(
		&connectBtn, r,
	)

	return connectBtn.btn
}

func Request() gui.Page {
	request := gui.CreatePage("Request")
	infoText := widget.NewLabel("Not connected")
	inputText := widget.NewMultiLineEntry()
	reqUrl := widget.NewEntry()
	connectBtn := connectionButton(infoText, reqUrl)

	inputText.SetPlaceHolder("Some request info...")
	reqUrl.SetPlaceHolder("wss://example.com")

	// REMOVE
	reqUrl.SetText("ws://localhost:8080")

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
