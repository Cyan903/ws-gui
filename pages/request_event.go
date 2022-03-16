package pages

import (
	"net/url"

	"fyne.io/fyne/v2/widget"
	"github.com/Cyan903/ws-gui/client"
)

type conButton struct {
	connected bool
	con       client.Client
	lbl       *widget.Label
	btn       *widget.Button
}

func (c *conButton) Connect(nurl string) {
	u, err := url.Parse(nurl)

	if err != nil {
		c.lbl.SetText("Invalid url")
		return
	}

	c.lbl.SetText("Awaiting connection...")
	c.con.Url = *u
	err = c.con.Connect()

	if err != nil {
		c.lbl.SetText("Error connecting!")
		return
	}

	c.lbl.SetText("Connected!")
	c.btn.SetText("Disconnect")
	c.connected = true
}

func (c *conButton) Disconnect(err error) {
	c.btn.SetText("Connect")

	if err != nil {
		c.lbl.SetText(err.Error())
	}

	if c.connected {
		c.con.Close()
	}

	c.connected = false
}

func toggleConnect(c *conButton, r *widget.Entry) func() {
	return func() {
		if !c.connected {
			c.Connect(r.Text)
			return
		}

		c.Disconnect(nil)
		c.lbl.SetText("Disconnected")
	}
}
