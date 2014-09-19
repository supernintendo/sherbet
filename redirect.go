package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func frameHandler(w http.ResponseWriter, r *http.Request) {
	response, err := http.Get(conf.Frame + r.URL.Path[1:])

	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		}

		w.Header().Set("Content-Type", string(response.Header.Get("Content-Type")))
		w.Write(contents)
	}
}
