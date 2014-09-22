package main

func dealWith(msg []byte) {
	switch string(msg) {
	case "connected":
		sendAll("status", []byte("Connected"))
	}
}
