package client

import (
	"github.com/gorilla/websocket"
)

func (c *Client) Send(msg string) {
	err := c.Con.WriteMessage(websocket.TextMessage, []byte(msg))

	if err != nil {
		return
	}
}

func (c *Client) Close() {
	err := c.Con.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))

	if err != nil {
		return
	}

	c.Con.Close()
}

func (c *Client) Connect() error {
	con, _, err := websocket.DefaultDialer.Dial(c.Url.String(), nil)
	c.Con = con

	if err != nil {
		return err
	}

	go func() {
		for {
			_, message, err := c.Con.ReadMessage()
			c.OnMessage(message, err)

			if err != nil {
				c.OnError(err)
				c.Close()

				return
			}
		}
	}()

	return nil
}
