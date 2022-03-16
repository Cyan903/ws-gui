package client

import (
	"net/url"

	"github.com/gorilla/websocket"
)

type Client struct {
	Url       url.URL
	Con       *websocket.Conn
	OnMessage func([]byte, error)
}
