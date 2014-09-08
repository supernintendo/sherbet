package main

import (
	"code.google.com/p/go.exp/fsnotify"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func newWatcher(path string, css string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			select {
			case ev := <-watcher.Event:
				if ev.IsModify() {
					s := strings.Split(ev.Name, "/")

					if s[len(s)-1] == css {
						file, err := ioutil.ReadFile(ev.Name)
						if err != nil {
							fmt.Print("Error reading file.", err)
						}
						o := append([]byte("css:sherbet###"), file...)

						sendAll([]byte(o))
					}
				}
			case err := <-watcher.Error:
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Watch(path)
	if err != nil {
		log.Fatal(err)
	}
}
