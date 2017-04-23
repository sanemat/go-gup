package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os/exec"

	"github.com/sanemat/go-gup/script/ghgutils"
	"github.com/sanemat/go-gup/script/gitdescribetags"
	"github.com/sanemat/go-gup/script/gogetutils"
)

func runGhr(pre bool) {
	ghrPath, err := ghgutils.GhgLookOrGet("ghr", "tcnksm/ghr")
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
	if err := gogetutils.GoGet("github.com/Songmu/ghg/cmd/ghg"); err != nil {
		log.Fatal(err)
	}
	if err := ghgutils.GhgGet("tcnksm/ghr"); err != nil {
		log.Fatal(err)
	}
	runGhr(pre)
}
