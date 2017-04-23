package main

import (
	"bytes"
	"fmt"
	"github.com/sanemat/go-gup/script/ghgutils"
	"github.com/sanemat/go-gup/script/gitdescribetags"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func rmRfPkg() error {
	if err := os.RemoveAll("pkg"); err != nil {
		return err
	}
	return nil
}

func goxRun() error {
	goxPath, err := ghgutils.GhgLookOrGetGox()
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
		log.Panic(err)
		log.Fatal(err)
	}
	if err := ghgutils.GoGetGhg(); err != nil {
		log.Panic(err)
		log.Fatal(err)
	}
	if err := ghgutils.GhgGet("mitchellh/gox"); err != nil {
		log.Panic(err)
		log.Fatal(err)
	}
	if err := goxRun(); err != nil {
		log.Panic(err)
		log.Fatal(err)
	}
}
