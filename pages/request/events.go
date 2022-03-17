package request

import (
	"net/url"
	"time"

	"fyne.io/fyne/v2/widget"
	"github.com/Cyan903/ws-gui/client"
	"github.com/Cyan903/ws-gui/pages/history"
	"github.com/Cyan903/ws-gui/pages/response"
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
	response.Connection("Connected to", nurl)
	history.HistoryData.Add([4]string{"",
		time.Now().Format(time.Kitchen),
		u.Scheme + "://",
		u.Host,
	})
	
	c.connected = true
}

func (c *conButton) Disconnect(err error) {
	c.btn.SetText("Connect")

	if err != nil {
		c.lbl.SetText(err.Error())

		if c.connected {
			response.Print("error", err.Error())
		}
	}

	if c.connected {
		response.Connection("Disconnected from", c.con.Url.String())
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
