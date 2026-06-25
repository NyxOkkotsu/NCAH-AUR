package main

import (
	"flag"
	"fmt"
	"os"

	"ncah/internal/aur"
	"ncah/internal/build"
	"ncah/internal/security"
	"ncah/internal/ui"
)

func main() {
	ui.PrintBanner()

	searchFlag := flag.String("Ss", "", "Search for super-cute AUR packages meww!")
	infoFlag := flag.String("Si", "", "Peek at the package info details nyaa~")
	installFlag := flag.String("S", "", "Install a package with maximum love and safety!")

	flag.Parse()

	if *searchFlag != "" {
		handleSearch(*searchFlag)
		return
	}

	if *infoFlag != "" {
		handleInfo(*infoFlag)
		return
	}

	if *installFlag != "" {
		handleInstall(*installFlag)
		return
	}

	fmt.Println("👉 Use -Ss [query] to search, -Si [pkg] to peek, or -S [pkg] to install nyaa~ (✿•ᴗ•)")
}

func handleSearch(query string) {
	results, err := aur.SearchPackages(query)
	if err != nil {
		ui.PrintFailure()
		return
	}
	ui.PrintSearchResults(results)
}

func handleInfo(pkgName string) {
	info, err := aur.GetPackageInfo(pkgName)
	if err != nil || info == nil {
		ui.PrintFailure()
		return
	}
	ui.PrintPackageInfo(info)
}

func handleInstall(pkgName string) {
	info, err := aur.GetPackageInfo(pkgName)
	if err != nil || info == nil {
		ui.PrintFailure()
		return
	}

	// 1. Dependency Transparency
	fmt.Printf("\n✨ \033[1;35m[Found some buddies needed for %s nyaww~]\033[0m\n", pkgName)
	ui.PrintDependencies(info.Depends, info.MakeDepends)
	if !ui.AskConfirmation("Proceed with installation?") {
		fmt.Println("Gumuuu... Installation cancelled by user! (っ°´o`°ς)")
		return
	}

	// 2. Fetch PKGBUILD
	fmt.Println("🛸 Fetching PKGBUILD from AUR... please wait a lil bit meww...")
	pkgbuildContent, dirPath, err := aur.FetchPKGBUILD(pkgName, info.PackageBase)
	if err != nil {
		ui.PrintFailure()
		return
	}
	defer os.RemoveAll(dirPath)

	// 3. Security Inspection
	fmt.Println("\n🔎 Nyx is closely inspecting the PKGBUILD contents right now...")
	ui.PrintHorizontalRule()
	fmt.Println(pkgbuildContent)
	ui.PrintHorizontalRule()

	report := security.ScanPKGBUILD(pkgbuildContent)
	ui.PrintSecurityReport(report)

	if report.RiskLevel == security.HighRisk {
		if !ui.AskConfirmation("⚠️ WARNING: This package is HIGH RISK! Are you absolutely sure you want to proceed?") {
			fmt.Println("Phew~ Smart choice! Safety first nyaww! 🛡️( •̀ ω •́ )✧")
			return
		}
	}

	// 4. Sudo Safety Prompt
	if ui.AskConfirmation("Wanna Install This AUR as sudo?") {
		ui.PrintSudoWarning()
	}

	// 5. Build and Install
	fmt.Println("🛠️  Starting the magic build process with makepkg...")
	err = build.BuildAndInstall(dirPath)
	if err != nil {
		ui.PrintFailure()
		return
	}

	// 6. Optional Dependency Cleanup
	if ui.AskConfirmation("Remove unused dependencies after install?") {
		fmt.Println("🧹 Sweeping away unused dependencies like dust...")
		_ = build.CleanupDependencies()
	}

	ui.PrintSuccess()
}