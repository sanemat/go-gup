package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/sanemat/go-gup/script/gitdescribetags"
	"github.com/sanemat/go-gup/script/gogetutils"
)

func rmRfPkg() error {
	version, err := gitdescribetags.Get()
	if err != nil {
		return err
	}
	if err := os.RemoveAll(filepath.Join("pkg", version)); err != nil {
		return err
	}
	return nil
}

func goxRun() error {
	goxPath, err := gogetutils.LookOrInstall("gox", "github.com/mitchellh/gox")
	if err != nil {
		return err
	}
	version, err := gitdescribetags.Get()
	if err != nil {
		return err
	}
	cmd := exec.Command(
		goxPath,
		"-output",
		filepath.Join("pkg", version, "{{.Dir}}_"+version+"_{{.OS}}_{{.Arch}}", "{{.Dir}}"),
		"-ldflags",
		"-X main.version="+version,
		"./cmd/gup",
	)
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return err
	}
	fmt.Print(out.String())
	return nil
}

func main() {
	if err := rmRfPkg(); err != nil {
		log.Fatal(err)
	}
	if _, err := gogetutils.LookOrInstall("gox", "github.com/mitchellh/gox"); err != nil {
		log.Fatal(err)
	}
	if err := goxRun(); err != nil {
		log.Fatal(err)
	}
}
