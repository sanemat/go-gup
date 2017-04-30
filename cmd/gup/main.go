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
	var protocol string

	flag.BoolVar(&showVersion, "v", false, "show version")
	flag.BoolVar(&showVersion, "version", false, "show version")
	flag.StringVar(&protocol, "protocol", "http", "protocol")
	flag.Parse()
	if showVersion {
		fmt.Println(version)
		return
	}
	fmt.Printf("* Listening on %s://%s%s\n", protocol, listeningHost, listeningAddr)
	fmt.Print("Use Ctrl-C to stop\n")
	http.HandleFunc("/", handler)
	if protocol == "https" {
		//# Key considerations for algorithm "RSA" ≥ 2048-bit
		//openssl genrsa -out server.key 2048
		//
		//# Key considerations for algorithm "ECDSA" ≥ secp384r1
		//# List ECDSA the supported curves (openssl ecparam -list_curves)
		//openssl ecparam -genkey -name secp384r1 -out server.key
		//
		//openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650  -subj "/CN=localhost"
		http.ListenAndServeTLS(listeningAddr, "server.crt", "server.key", Log(http.DefaultServeMux))
	} else {
		http.ListenAndServe(listeningAddr, Log(http.DefaultServeMux))
	}
}
