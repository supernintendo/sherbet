package main

// import (
// 	"io/ioutil"
// )

func dealWith(msg []byte) {
	switch string(msg) {
	case "Connected":
		sendAll("Status", make([]byte, 0, 1), "Connected.")
	}
}
