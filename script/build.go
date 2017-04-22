package main

import (
	"bytes"
	"log"
	"os/exec"
)

func goGetGox() {
	goPath, err := exec.LookPath("go")
	if err != nil {
		log.Fatal(err)
	}
	cmd := exec.Command(goPath, "get", "github.com/mitchellh/gox")
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}

func goxRun() {
	goxPath, err := exec.LookPath("gox")
	if err != nil {
		log.Fatal(err)
	}
	cmd := exec.Command(
		goxPath,
		"-output",
		"pkg/{{.OS}}_{{.Arch}}/{{.Dir}}",
		"./cmd/gup",
	)
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	goGetGox()
	goxRun()
}
