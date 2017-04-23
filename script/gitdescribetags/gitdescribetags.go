package gitdescribetags

import (
	"os/exec"
	"bytes"
	"strings"
)

func Get() (string, error) {
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
	return strings.TrimSpace(out.String()), nil
}
