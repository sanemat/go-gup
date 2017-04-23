package ghgutils

import (
	"bytes"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func GhgGet(target string) error {
	ghgPath, err := exec.LookPath("ghg")
	if err != nil {
		log.Panic(err)
		return err
	}
	cmd := exec.Command(ghgPath, "get", target)
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		log.Panic(err)
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
		if err := GhgGet("tcnksm/ghr"); err != nil {
			return "", err
		}
		if _, err := os.Stat(ghrPath); os.IsNotExist(err) {
			return "", err
		}
	}
	return ghrPath, nil
}

func GhgLookOrGetGox() (string, error) {
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
	goxPath := filepath.Join(strings.TrimSpace(out.String()), "gox")
	if _, err := os.Stat(goxPath); os.IsNotExist(err) {
		if err := GhgGet("mitchellh/gox"); err != nil {
			return "", err
		}
		if _, err := os.Stat(goxPath); os.IsNotExist(err) {
			return "", err
		}
	}
	return goxPath, nil
}
