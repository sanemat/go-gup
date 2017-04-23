package gogetutils

import (
	"bytes"
	"os/exec"
)

func GoGet(packagePath string) error {
	goPath, err := exec.LookPath("go")
	if err != nil {
		return err
	}
	cmd := exec.Command(goPath, "get", packagePath)
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func LookOrInstall(commandName string, packagePath string) (string, error) {
	targetPath, err := exec.LookPath(commandName)
	if err != nil {
		if err := GoGet(packagePath); err != nil {
			return "", err
		} else {
			targetPath, err = exec.LookPath(commandName)
			if err != nil {
				return "", err
			}
		}
	}
	return targetPath, nil
}
