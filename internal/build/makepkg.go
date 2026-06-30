package build

import (
	"os"
	"os/exec"
	"strings"
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

func ScanMalware(content string) []string {
	var findings []string
	badPatterns := []string{
		"curl ", "wget ", "rm -rf /", "base64 ", "nc ", "netcat ",
		"/dev/tcp", "socat ", "crontab ", "eval ", "raw.githubusercontent",
	}
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "#") {
			continue
		}
		for _, pattern := range badPatterns {
			if strings.Contains(line, pattern) {
				findings = append(findings, trimmed)
				break
			}
		}
	}
	return findings
}

func ScanVirus(dirPath string) (string, bool) {
	_, err := exec.LookPath("clamscan")
	if err != nil {
		cmdPkg := exec.Command("sudo", "pacman", "-S", "clamav", "--noconfirm")
		cmdPkg.Stdout = os.Stdout
		cmdPkg.Stderr = os.Stderr
		if err := cmdPkg.Run(); err != nil {
			return "Failed to automatically install ClamAV dependency via pacman.", false
		}
		cmdDb := exec.Command("sudo", "freshclam")
		cmdDb.Stdout = os.Stdout
		cmdDb.Stderr = os.Stderr
		if err := cmdDb.Run(); err != nil {
			return "ClamAV installed, but failed to sync database via freshclam.", false
		}
	}
	cmd := exec.Command("clamscan", "-r", "--no-summary", dirPath)
	out, _ := cmd.CombinedOutput()
	if cmd.ProcessState != nil && cmd.ProcessState.ExitCode() == 1 {
		return string(out), true
	}
	return "", false
}

func EditPKGBUILD(dirPath string) error {
	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = "nano"
	}
	cmd := exec.Command(editor, dirPath+"/PKGBUILD")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func SyncOfficial(flag string) error {
	cmd := exec.Command("sudo", "pacman", flag)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func GetForeignPackages() (map[string]string, error) {
	cmd := exec.Command("pacman", "-Qm")
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	pkgs := make(map[string]string)
	lines := strings.Split(string(out), "\n")
	for _, line := range lines {
		parts := strings.Fields(line)
		if len(parts) == 2 {
			pkgs[parts[0]] = parts[1]
		}
	}
	return pkgs, nil
}

func CheckVersionUpgrade(local, remote string) bool {
	cmd := exec.Command("vercmp", local, remote)
	out, err := cmd.Output()
	if err != nil {
		return false
	}
	res := strings.TrimSpace(string(out))
	return strings.HasPrefix(res, "-")
}