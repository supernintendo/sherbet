package main

import (
	"github.com/gorilla/websocket"
)

func dealWith(msg []byte, c *websocket.Conn) {
	switch string(msg) {
	case "Connected":
		sendTo(c, "Status", make([]byte, 0, 1), "Connected.")
	}
}
