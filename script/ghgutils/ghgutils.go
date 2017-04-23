package ghgutils

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func GoGetGhg() error {
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
func LookOrInstall() (string, error) {
	ghgPath, err := exec.LookPath("ghg")
	if err != nil {
		if err := GoGetGhg(); err != nil {
			return "", err
		} else {
			ghgPath, err = exec.LookPath("ghg")
			if err != nil {
				return "", err
			}
		}
	}
	return ghgPath, nil
}
func GhgGetGhr() error {
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
func GhgLookOrGetGhr() (string, error) {
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
	ghrPath := filepath.Join(strings.TrimSpace(out.String()), "ghr")
	if _, err := os.Stat(ghrPath); os.IsNotExist(err) {
		if err := GhgGetGhr(); err != nil {
			return "", err
		}
		if _, err := os.Stat(ghrPath); os.IsNotExist(err) {
			return "", err
		}
	}
	return ghrPath, nil
}
