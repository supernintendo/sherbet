package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
)

type Sherbetfile struct {
	Port                       int
	Frame, ReloadCSS, Watch    string
}

var conf Sherbetfile
var framehost string

func main() {
	// Read sherbet.json
	jsonFile := flag.String("json", "./sherbet.json", "json file to use")
	flag.Parse()
	content, err := ioutil.ReadFile(*jsonFile)
	rootPath := "./" + filepath.Dir(*jsonFile)
	if err != nil {
		fmt.Print("Error reading sherbet.json.", err)
	}
	err = json.Unmarshal(content, &conf)
	if err != nil {
		fmt.Print("Error parsing sherbet.json.", err)
	}
	parseFrameString()
	setupServer(conf.Port, conf.ReloadCSS, rootPath+"/"+conf.Watch)
}

func parseFrameString() {
	if string(conf.Frame[len(conf.Frame) - 1]) == "/" {
		framehost = conf.Frame
	} else {
		framehost = conf.Frame + "/"
	}
}

func setupServer(port int, css string, watch string) {
	// Make a new slice for our WebSocket connections.
	connections = make(map[*websocket.Conn]bool)
	router := mux.NewRouter()
	router.Host("sherbet").HandlerFunc(frameHandler)
	router.HandleFunc("/ws", wsHandler)
	http.Handle("/", router)

	// Watch the directory for changes, sending a socket message if a change
	// was made to the CSS.
	newWatcher(watch, css)

	log.Printf("Running on port %d\n", port)
	addr := fmt.Sprintf("127.0.0.1:%d", port)
	err := http.ListenAndServe(addr, nil)
	fmt.Println(err.Error())
}
