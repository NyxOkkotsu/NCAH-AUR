package aur

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
)

func FetchPKGBUILD(pkgName string, pkgBase string) (string, string, error) {
	tmpDir, err := os.MkdirTemp("", "ncah-*")
	if err != nil {
		return "", "", err
	}

	cloneUrl := fmt.Sprintf("https://aur.archlinux.org/%s.git", pkgBase)
	cmd := exec.Command("git", "clone", cloneUrl, tmpDir)
	if err := cmd.Run(); err != nil {
		os.RemoveAll(tmpDir)
		return "", "", err
	}

	pkgbuildPath := filepath.Join(tmpDir, "PKGBUILD")
	file, err := os.Open(pkgbuildPath)
	if err != nil {
		os.RemoveAll(tmpDir)
		return "", "", err
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		os.RemoveAll(tmpDir)
		return "", "", err
	}

	return string(content), tmpDir, nil
}