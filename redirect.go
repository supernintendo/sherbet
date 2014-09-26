package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func frameHandler(w http.ResponseWriter, r *http.Request) {
	transport := &http.Transport{DisableCompression: true}
	client := &http.Client{Transport: transport}

	response, err := client.Get(app.Frame + r.URL.Path[1:])

	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)

		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		} else {
			w.Header().Set("Content-Type", string(response.Header.Get("Content-Type")))

			if strings.Contains(string(response.Header.Get("Content-Type")), "text/html") {
				s := strings.Split(string(contents), "<head>")
				javascript, err := build_bundle_js()

				if err != nil {
					fmt.Printf("%s", err)
					os.Exit(1)
				} else {
					top, bottom := s[0], s[1]
					a := append([]byte(top), []byte("<head><script>")...)
					b := append(a, javascript...)
					c := append(b, []byte("</script>")...)
					finalContents := append(c, []byte(bottom)...)

					w.Write(finalContents)
				}
			} else {
				w.WriteHeader(200)
				w.Write(contents)
			}
		}
	}
}
