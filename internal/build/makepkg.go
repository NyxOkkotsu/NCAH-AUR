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
	return cmd.Run()
}

func CleanupDependencies() error {
	cmd := exec.Command("bash", "-c", "sudo pacman -Rns $(pacman -Qtdq) --noconfirm")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func RemovePackage(pkgName string) error {
	cmd := exec.Command("sudo", "pacman", "-Rns", pkgName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}