package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"github.com/sanemat/go-gup/script/gitdescribetags"
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
	version, err := gitdescribetags.Get()
	if err != nil {
		log.Fatal(err)
	}
	cmd := exec.Command(
		goxPath,
		"-output",
		filepath.Join("pkg", version, "{{.Dir}}_" + version+ "_{{.OS}}_{{.Arch}}", "{{.Dir}}"),
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
