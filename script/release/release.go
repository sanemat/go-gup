package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os/exec"

	"os"

	"github.com/sanemat/go-gup/script/ghgutils"
	"github.com/sanemat/go-gup/script/gitdescribetags"
	"github.com/sanemat/go-gup/script/gogetutils"
)

const (
	EnvCI             = "CI"
	EnvCircleCi       = "CIRCLECI"
	EnvCircleCiForTag = "CIRCLE_TAG"
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

	params := []string{
		"-u",
		"sanemat",
		"-r",
		"go-gup",
	}
	if pre {
		params = append(params, "-prerelease")
	}
	params = append(params, version, "./pkg/")
	cmd := exec.Command(ghrPath, params...)
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		log.Print(out.String())
		log.Fatal(err)
	}
	fmt.Print(out.String())
}
func circleCIBuildForTag() (bool, error) {
	gitTag := os.Getenv(EnvCircleCiForTag)
	buildForTag := gitTag != nil && gitTag != ""
	return buildForTag, nil
}

func main() {
	var pre bool
	flag.BoolVar(&pre, "pre", false, "pre release")
	flag.Parse()

	circleCiBuildForTag, err := circleCIBuildForTag()
	if err != nil {
		log.Fatal(err)
	}
	if pre && circleCiBuildForTag {
		return
	}

	if err := gogetutils.GoGet("github.com/Songmu/ghg/cmd/ghg"); err != nil {
		log.Fatal(err)
	}
	if err := ghgutils.GhgGet("tcnksm/ghr"); err != nil {
		log.Fatal(err)
	}
	runGhr(pre)
}
