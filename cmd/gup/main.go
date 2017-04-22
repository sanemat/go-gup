package main

import (
	"flag"
	"fmt"
	"net/http"
)

var version = "0.1.0"

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	var showVersion bool
	var listeningAddr = ":8080"
	var listeningHost = "localhost"

	flag.BoolVar(&showVersion, "v", false, "show version")
	flag.BoolVar(&showVersion, "version", false, "show version")
	flag.Parse()
	if showVersion {
		fmt.Println(version)
		return
	}
	fmt.Printf("* Listening on http://%s%s\n", listeningHost, listeningAddr)
	fmt.Print("Use Ctrl-C to stop\n")
	http.HandleFunc("/", handler)
	http.ListenAndServe(listeningAddr, nil)
}
