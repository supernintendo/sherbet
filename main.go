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
	Directory                  string
	Frame                      string
	Port                       int
	RootPath                   string
	Target                     string
	Watchers                   []Watcher
}

func (s *Sherbetfile) SetFrame(f string) string {
	if s.Target[len(s.Target) - 1] == '/' {
		s.Frame = f
	} else {
		s.Frame = f + "/"
	}
	return s.Frame
}

var app Sherbetfile

func main() {
	// Read sherbet.json
	jsonFile := flag.String("j", "./sherbet.json", "Sherbet JSON configuration file.")
	flag.Parse()

	s, err := ioutil.ReadFile(*jsonFile)

	if err != nil {
		fmt.Print("Error reading sherbet.json.", err)
	}
	err = json.Unmarshal(s, &app)

	if err != nil {
		fmt.Print("Error parsing sherbet.json.", err)
	}
	app.RootPath = "./" + filepath.Dir(*jsonFile)
	app.SetFrame(app.Target)
	setupServer()
}

// func setupServer(port int, css string, watch string) {
func setupServer() {
	// Make a new slice for our WebSocket connections.
	connections = make(map[*websocket.Conn]bool)
	router := mux.NewRouter()
	router.Host("sherbet").HandlerFunc(frameHandler)
	router.HandleFunc("/ws", wsHandler)
	http.Handle("/", router)

	// Watch the directory for changes, sending a socket message if a change
	// was made to the CSS.

	for i := range app.Watchers {
		activateWatcher(app.Watchers[i])
	}

	log.Printf("Running on port %d\n", app.Port)
	addr := fmt.Sprintf("127.0.0.1:%d", app.Port)
	err := http.ListenAndServe(addr, nil)
	fmt.Println(err.Error())
}
