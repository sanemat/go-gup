package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func getVersion2() (string, error) {
	cmd := exec.Command(
		"git",
		"describe",
		"--tags",
	)
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return "", err
	}
	return out.String(), nil
}

func goGetGhg() error {
	goPath, err := exec.LookPath("go")
	if err != nil {
		return err
	}
	cmd := exec.Command(goPath, "get", "github.com/Songmu/ghg/cmd/ghg")
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
func ghgGetGhr() error {
	ghgPath, err := exec.LookPath("ghg")
	if err != nil {
		return err
	}
	cmd := exec.Command(ghgPath, "get", "tcnksm/ghr")
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func getGhrPath() (string, error) {
	//$(ghg bin)/ghr
	ghgPath, err := exec.LookPath("ghg")
	if err != nil {
		return "", err
	}
	cmd := exec.Command(ghgPath, "bin")
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return "", err
	}
	return filepath.Join(out.String(), "ghr"), nil
}
func runGhr(pre bool) error {
	ghrPath, err := getGhrPath()
	if err != nil {
		return err
	}
	version, err := getVersion2()
	if err != nil {
		return err
	}

	var cmd exec.Cmd
	if pre {
		cmd = exec.Command(
			ghrPath,
			"-prerelease",
			version,
			"pkg/",
		)
	} else {
		cmd = exec.Command(
			ghrPath,
			version,
			"pkg/",
		)
	}
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return err
	}
	fmt.Print(out.String())
	return nil
}

func main() {
	var pre bool
	flag.BoolVar(&pre, "pre", false, "pre release")
	flag.Parse()
	if err := goGetGhg(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	if err := ghgGetGhr(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	if err := runGhr(pre); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
