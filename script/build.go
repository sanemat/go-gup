package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func getVersion1() (string, error) {
	cmd := exec.Command(
		"git",
		"describe",
		"--tags",
	)
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
		return "", err
	}
	return out.String(), nil
}

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

func rmRfPkg() {
	if err := os.RemoveAll("pkg"); err != nil {
		log.Fatal(err)
	}
}

func goxRun() {
	goxPath, err := exec.LookPath("gox")
	if err != nil {
		log.Fatal(err)
	}
	version, err := getVersion1()
	if err != nil {
		log.Fatal(err)
	}
	cmd := exec.Command(
		goxPath,
		"-output",
		"pkg/{{.OS}}_{{.Arch}}/{{.Dir}}",
		"-ldflags",
		"-X main.version="+version,
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
	rmRfPkg()
	goGetGox()
	goxRun()
}
