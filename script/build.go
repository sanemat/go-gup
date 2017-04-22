package main

import (
	"bytes"
	"log"
	"os/exec"
	"fmt"
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

func getVersion() string {
	cmd := exec.Command(
		"git",
		"describe",
		"--tags",
	)
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
	return out.String();
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
		"-ldflags",
		"-X main.version=" + getVersion(),
		"./cmd/gup",
	)
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
	fmt.Print(out.String())
}

func main() {
	goGetGox()
	goxRun()
}
