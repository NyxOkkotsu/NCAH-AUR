package ui

import (
	"fmt"
	"strings"
	"ncah/internal/aur"
	"ncah/internal/security"
)

func PrintHorizontalRule() {
	fmt.Println("\033[1;30m~🐾~🐾~🐾~🐾~🐾~🐾~🐾~🐾~🐾~🐾~🐾~🐾~🐾~🐾~🐾~🐾~🐾~\033[0m")
}

func PrintFailure() {
	fmt.Println("\n\033[1;31mEhh... !! sorry~~ The AUR isn't installed, it is my mistakes? ｡°(°.◜ᯅ◝°)°｡\033[0m")
}

func PrintSuccess() {
	fmt.Println("\n\033[1;32mThe AUR succesfully Installed nyaww:3 🎉✨ (っ^‿^)っ\033[0m")
}

func PrintSudoWarning() {
	fmt.Println("\033[1;33m" + `[!] "This package requires root privileges to proceed."` + "\033[0m")
}

func PrintSearchResults(results []aur.AURPackage) {
	if len(results) == 0 {
		fmt.Println("Gomennasai~ Couldn't find any cute packages matching that (っ- ‸ -,)\033[0m")
		return
	}
	fmt.Println("🐾 Look what Nyx found for you:")
	for _, pkg := range results {
		fmt.Printf("\033[1;36maur/\033[1;37m%s \033[1;32m%s\033[0m\n    💌 %s\n", pkg.Name, pkg.Version, pkg.Description)
	}
}

func PrintPackageInfo(pkg *aur.AURPackage) {
	fmt.Printf("\033[1;35mPackage Name :\033[0m %s 🐾\n", pkg.Name)
	fmt.Printf("\033[1;35mCute Version :\033[0m %s\n", pkg.Version)
	fmt.Printf("\033[1;35mDescription  :\033[0m %s\n", pkg.Description)
	fmt.Printf("\033[1;35mHome Link    :\033[0m %s\n", pkg.URL)
	fmt.Printf("\033[1;35mCare Taker   :\033[0m %s\n", pkg.Maintainer)
}

func PrintDependencies(depends []string, makeDepends []string) {
	if len(depends) > 0 {
		fmt.Printf("📦 Core Buddies: %s\n", stringify(depends))
	}
	if len(makeDepends) > 0 {
		fmt.Printf("🛠️  Build Buddies: %s\n", stringify(makeDepends))
	}
	if len(depends) == 0 && len(makeDepends) == 0 {
		fmt.Println("Woww! This package is super independent, no buddies needed meww~")
	}
}

func PrintSecurityReport(report *security.ScanReport) {
	var color string
	switch report.RiskLevel {
	case security.Safe:
		color = "\033[1;32m[SAFE] ✨ Nyx loves this pack!\033[0m"
	case security.Warning:
		color = "\033[1;33m[WARNING] ⚠️ Please follow your heart carefully~\033[0m"
	case security.HighRisk:
		color = "\033[1;31m[HIGH RISK] 🔥🙀 Oh noes, super scary!\033[0m"
	}

	fmt.Printf("\n🛡️  \033[1;37mNyx's Security X-Ray Result:\033[0m %s\n", color)
	if len(report.Findings) > 0 {
		fmt.Println("🚨 Uh-oh, found some naughty bits inside:")
		for _, finding := range report.Findings {
			fmt.Printf("   - %s\n", finding)
		}
	} else {
		fmt.Println("   Squeaky clean! Everything looks wholesome and safe nyaa~")
	}
	fmt.Println()
}

func stringify(slice []string) string {
	if len(slice) == 0 {
		return "None meww"
	}
	return "[" + strings.Join(slice, ", ") + "]"
}