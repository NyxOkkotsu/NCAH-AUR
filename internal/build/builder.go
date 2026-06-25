package build

import (
	"os"
	"os/exec"
)

func BuildAndInstall(dirPath string) error {
	cmd := exec.Command("makepkg", "-si", "--noconfirm")
	cmd.Dir = dirPath
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	return cmd.Run()
}