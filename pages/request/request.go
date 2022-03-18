package request

import (
	"net/url"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Cyan903/ws-gui/client"
	"github.com/Cyan903/ws-gui/gui"
	"github.com/Cyan903/ws-gui/pages/response"
)

func connectionButton(t *widget.Label, r *widget.Entry) *conButton {
	connectBtn := conButton{}
	client := client.Client{
		Url: url.URL{},

		OnMessage: func(msg []byte, err error) {
			if err != nil {
				return
			}

			response.Print("log", string(msg))
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

	return &connectBtn
}

func Request() gui.Page {
	request := gui.CreatePage("Request")
	infoText := widget.NewLabel("Not connected")
	inputText := widget.NewMultiLineEntry()
	reqUrl := widget.NewEntry()
	connectBtn := connectionButton(infoText, reqUrl)

	inputText.SetPlaceHolder("Some request info...")
	reqUrl.SetPlaceHolder("wss://example.com")

	inpForm := &widget.Form{
		Items: []*widget.FormItem{{
			Widget: container.NewAdaptiveGrid(2,
				connectBtn.btn, container.NewCenter(infoText),
			),
		}, {
			Widget: reqUrl,
		}, {
			Widget: inputText,
		}},

		OnSubmit: func() { sendMessage(connectBtn, inputText.Text) },
	}

	inpForm.SubmitText = "Send"
	request.AddItem(inpForm)

	return request
}
