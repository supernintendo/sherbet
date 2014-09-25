package main

import (
	"code.google.com/p/go.exp/fsnotify"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type Watcher struct {
	Category     string
	Create       bool
	DeliverFile  bool
	File         string
	Message      string
	OnCreate     bool
	OnDelete     bool
	OnLoad       bool
	OnModify     bool
	OnRename     bool
}

func activateWatcher(w Watcher) {
	activated, err := fsnotify.NewWatcher()

	if err != nil {
		log.Fatal(err)
	}
	go func() {
		for {
			select {
			case event := <-activated.Event:
				slugs := strings.Split(event.Name, "/")

				if slugs[len(slugs)-1] == w.File {

					c := event.IsCreate() && w.OnCreate
					d := event.IsDelete() && w.OnDelete
					m := event.IsModify() && w.OnModify
					r := event.IsRename() && w.OnRename

					file, err := ioutil.ReadFile(event.Name)

					if err != nil {
						fmt.Print("Error reading file.", err)
					}
					switch {
					case c: sendAll(w.Category, make([]byte, 0, 1), w.Message)
					case d: sendAll(w.Category, make([]byte, 0, 1), w.Message)
					case m: sendAll(w.Category, file, w.Message)
					case r: sendAll(w.Category, file, w.Message)
					}
				}
			case err := <-activated.Error:
				log.Println("error:", err)
			}
		}
	}()
	err = activated.Watch(sherbetFile.RootPath)
	if err != nil {
		log.Fatal(err)
	}
}
