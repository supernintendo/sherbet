package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var connections map[*websocket.Conn]bool

func sendAll(msg []byte) {
	for conn := range connections {
		if err := conn.WriteMessage(websocket.TextMessage, msg); err != nil {
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
