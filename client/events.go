package client

import (
	"log"

	"github.com/gorilla/websocket"
)

func (c *Client) Send(msg string) {
	err := c.Con.WriteMessage(websocket.TextMessage, []byte(msg))

	if err != nil {
		log.Println("write:", err)
		return
	}
}

func (c *Client) Close() {
	err := c.Con.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))

	if err != nil {
		log.Println("write close:", err)
		return
	}

	c.Con.Close()
}

func (c *Client) Connect() error {
	log.Printf("connecting to %s", c.Url.String())

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
