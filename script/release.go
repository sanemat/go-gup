package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/sanemat/go-gup/script/ghgutils"
	"github.com/sanemat/go-gup/script/gitdescribetags"
	"log"
	"os/exec"
)

func runGhr(pre bool) {
	ghrPath, err := ghgutils.GhgLookOrGetGhr()
	if err != nil {
		log.Fatal(err)
	}
	version, err := gitdescribetags.Get()
	if err != nil {
		log.Fatal(err)
	}

	var cmd *exec.Cmd
	if pre {
		cmd = exec.Command(
			ghrPath,
			"-prerelease",
			"-u",
			"sanemat",
			"-r",
			"go-gup",
			version,
			"./pkg/",
		)
	} else {
		cmd = exec.Command(
			ghrPath,
			"-u",
			"sanemat",
			"-r",
			"go-gup",
			version,
			"./pkg/",
		)
	}
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		log.Print(out.String())
		log.Fatal(err)
	}
	fmt.Print(out.String())
}

func main() {
	var pre bool
	flag.BoolVar(&pre, "pre", false, "pre release")
	flag.Parse()
	if err := ghgutils.GoGetGhg(); err != nil {
		log.Fatal(err)
	}
	if err := ghgutils.GhgGet("tcnksm/ghr"); err != nil {
		log.Fatal(err)
	}
	runGhr(pre)
}
