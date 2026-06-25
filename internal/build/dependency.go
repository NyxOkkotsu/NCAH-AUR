package build

import (
	"os"
	"os/exec"
)

func CleanupDependencies() error {
	cmd := exec.Command("sh", "-c", "sudo pacman -Rns $(pacman -Qtdq)")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}