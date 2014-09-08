package main

func dealWith(msg []byte) {
	switch string(msg) {
	case "connected":
		f := "frame:sherbet###" + conf.Frame

		sendAll([]byte(f))
	}
}
