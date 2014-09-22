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
			case event := <-watcher.Event:
				if event.IsModify() {
					fileSlugs := strings.Split(event.Name, "/")

					// Check that the event is related to the reloadCSS file.
					if fileSlugs[len(fileSlugs)-1] == css {
						file, err := ioutil.ReadFile(event.Name)
						if err != nil {
							fmt.Print("Error reading file.", err)
						}
						sendAll("reloadCSS", file)
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
