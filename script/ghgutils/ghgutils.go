package ghgutils

import (
	"bytes"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func GhgGet(githubUserRepo string) error {
	ghgPath, err := exec.LookPath("ghg")
	if err != nil {
		log.Panic(err)
		return err
	}
	cmd := exec.Command(ghgPath, "get", githubUserRepo)
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		log.Panic(err)
		return err
	}
	return nil
}
func GhgLookOrGet(commandName string, githubUserRepo string) (string, error) {
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
	targetPath := filepath.Join(strings.TrimSpace(out.String()), commandName)
	if _, err := os.Stat(targetPath); os.IsNotExist(err) {
		if err := GhgGet(githubUserRepo); err != nil {
			return "", err
		}
		if _, err := os.Stat(targetPath); os.IsNotExist(err) {
			return "", err
		}
	}
	return targetPath, nil
}
