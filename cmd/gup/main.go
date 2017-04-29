package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	log "github.com/Sirupsen/logrus"
)

var version = "v0.1.4"

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

func handler(w http.ResponseWriter, r *http.Request) {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	name := filepath.Join(cwd, r.URL.Path)
	f, err := os.Open(name)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	defer f.Close()
	io.Copy(w, f)
}

func Log(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.WithFields(log.Fields{
			"RemoteAddr": r.RemoteAddr,
			"Method": r.Method,
			"URL": r.URL,
			"Header": r.Header,
		}).Info("Access")

		handler.ServeHTTP(w, r)
	})
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
	http.ListenAndServe(listeningAddr, Log(http.DefaultServeMux))
}
