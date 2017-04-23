package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"log"
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
		return "", errors.New("It does not detect git describe.")
	}
	return out.String(), nil
}

func goGetGhg() {
	goPath, err := exec.LookPath("go")
	if err != nil {
		log.Fatal(err)
	}
	cmd := exec.Command(goPath, "get", "github.com/Songmu/ghg/cmd/ghg")
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
func ghgGetGhr() {
	ghgPath, err := exec.LookPath("ghg")
	if err != nil {
		log.Fatal(err)
	}
	cmd := exec.Command(ghgPath, "get", "tcnksm/ghr")
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}

func getGhrPath() string {
	//$(ghg bin)/ghr
	ghgPath, err := exec.LookPath("ghg")
	if err != nil {
		log.Fatal(err)
	}
	cmd := exec.Command(ghgPath, "bin")
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
	return filepath.Join(out.String(), "ghr")
}
func runGhr(pre bool) {
	ghrPath := getGhrPath()
	version, err := getVersion2()
	if err != nil {
		log.Fatal(err)
	}

	var cmd *exec.Cmd
	if pre {
		cmd = exec.Command(
			ghrPath,
			"-prerelease",
			"-debug",
			version,
			"pkg/",
		)
	} else {
		cmd = exec.Command(
			ghrPath,
			"-debug",
			version,
			"pkg/",
		)
	}
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		log.Print(out.String())
		log.Panic(err)
	}
	fmt.Print(out.String())
}

func main() {
	var pre bool
	flag.BoolVar(&pre, "pre", false, "pre release")
	flag.Parse()
	goGetGhg()
	ghgGetGhr()
	runGhr(pre)
}