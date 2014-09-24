package main

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type WSMessage struct {
    Category        string
	File            string
	Message         string
}

var connections map[*websocket.Conn]bool

func sendAll(c string, f []byte, m string) {
	msg := &WSMessage{Category: c, File: string(f), Message: m}
	output, err := json.Marshal(msg)

	if err != nil {
		log.Fatal(err)
	}
	for conn := range connections {
		if err := conn.WriteMessage(websocket.TextMessage, output); err != nil {
			delete(connections, conn)
			conn.Close()
		}
	}
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	// Upgrade the connection to a WebSocket connection
	conn, err := websocket.Upgrade(w, r, nil, 1024, 1024)
	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(w, "Not a websocket handshake", 400)
		return
	} else if err != nil {
		log.Println(err)
		return
	}
	connections[conn] = true

	for {
		// Blocks until a message is read
		_, msg, err := conn.ReadMessage()
		if err != nil {
			delete(connections, conn)
			conn.Close()
			return
		}
		dealWith(msg)
	}
}
