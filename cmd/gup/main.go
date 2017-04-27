package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var version = "v0.1.2"

func handler(w http.ResponseWriter, r *http.Request) {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	name := filepath.Join(cwd, filepath.Base(r.URL.Path))
	f, err := os.Open(name)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	defer f.Close()
	io.Copy(w, f)
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
