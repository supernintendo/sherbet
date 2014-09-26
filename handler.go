package main

import (
	"fmt"
	"io/ioutil"
	"github.com/gorilla/websocket"
)

func dealWith(msg []byte, c *websocket.Conn) {
	switch string(msg) {
	case "Connected":
		sendWatcherPayload(c)
	}
}

func sendWatcherPayload(c *websocket.Conn) {
	for _, w := range app.Watchers {
		if w.OnLoad {
			file, err := ioutil.ReadFile(app.RootPath + "/" + w.File)

			if err != nil {
				fmt.Print("Error reading file.", err)
			}
			sendTo(c, w.Category, file, w.Message)
		}
	}
}
