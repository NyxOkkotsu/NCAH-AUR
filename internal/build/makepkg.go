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
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func IsOfficialPackage(pkgName string) bool {
	cmd := exec.Command("pacman", "-Sp", pkgName)
	return cmd.Run() == nil
}

func InstallOfficialPackage(pkgName string) error {
	cmd := exec.Command("sudo", "pacman", "-S", pkgName)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func PrintOfficialSearch(query string) {
	cmd := exec.Command("pacman", "-Ss", query)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	_ = cmd.Run()
}

func PrintOfficialInfo(pkgName string) bool {
	cmd := exec.Command("pacman", "-Si", pkgName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run() == nil
}
