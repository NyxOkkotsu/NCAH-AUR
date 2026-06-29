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

	searchFlag := flag.String("Ss", "", "")
	infoFlag := flag.String("Si", "", "")
	installFlag := flag.String("S", "", "")
	removeFlag := flag.String("R", "", "")

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

	if *removeFlag != "" {
		handleRemove(*removeFlag)
		return
	}

	fmt.Println("👉 pssssst! use -Ss [query] to hunt, -Si [pkg] to peek, -S [pkg] to adopt, or -R [pkg] to banish a package nyaa~! (✿•ᴗ•)")
}

func handleSearch(query string) {
	ui.PrintStatus("sniffing official pacman databases... 🐾")
	build.PrintOfficialSearch(query)
	ui.PrintHorizontalRule()
	ui.PrintStatus("pouncing on the AUR database layout nyaa... 🐾")
	results, err := aur.SearchPackages(query)
	if err != nil {
		ui.PrintFailure()
		return
	}
	ui.PrintSearchResults(results)
}

func handleInfo(pkgName string) {
	ui.PrintStatus("peeking at official repository definitions... 📂")
	if build.PrintOfficialInfo(pkgName) {
		return
	}
	ui.PrintHorizontalRule()
	ui.PrintStatus("not found in official repos! tracking down AUR metadata uwu... 📂")
	info, err := aur.GetPackageInfo(pkgName)
	if err != nil || info == nil {
		ui.PrintFailure()
		return
	}
	ui.PrintPackageInfo(info)
}

func handleRemove(pkgName string) {
	ui.PrintStatus(fmt.Sprintf("preparing uninstallation protocol to banish %s nyaa...", pkgName))
	if !ui.AskConfirmation(fmt.Sprintf("Are you absolutely sure you want to purge %s and its configuration remnants from our cozy layout?", pkgName)) {
		ui.PrintCancel()
		return
	}
	err := build.RemovePackage(pkgName)
	if err != nil {
		ui.PrintCancel()
		return
	}
	fmt.Printf("\n\033[0;32m✔ %s has been successfully vaporized from the system layout! nyaa~\033[0m\n", pkgName)
}

func handleInstall(pkgName string) {
	if build.IsOfficialPackage(pkgName) {
		ui.PrintStatus("found inside official repositories! preparing native pacman deployment protocol nyaa~ ✨")
		ui.PrintSudoWarning()
		err := build.InstallOfficialPackage(pkgName)
		if err != nil {
			ui.PrintCancel()
			return
		}
		ui.PrintSuccess()
		return
	}

	ui.PrintStatus("waking up remote AUR api mirrors... mwah")
	info, err := aur.GetPackageInfo(pkgName)
	if err != nil || info == nil {
		ui.PrintFailure()
		return
	}

	fmt.Printf("\n✨ \033[1;36m[Dependencies validation target for: %s]\033[0m\n", pkgName)
	ui.PrintDependencies(info.Depends, info.MakeDepends)

	if !ui.AskConfirmation("Proceed with downloading installation recipes?") {
		ui.PrintCancel()
		return
	}

	ui.PrintStatus("fetching remote repository and caching staging area...")
	pkgbuildContent, dirPath, err := aur.FetchPKGBUILD(pkgName, info.PackageBase)
	if err != nil {
		ui.PrintFailure()
		return
	}
	defer os.RemoveAll(dirPath)

	pkgHash := security.CalculateHash(pkgbuildContent)
	isTrusted := security.IsHashTrusted(pkgHash)

	if isTrusted {
		ui.PrintHorizontalRule()
		fmt.Println("  \033[0;32m✔ Already reviewed & trusted by NCAH security core! Skipping strict verification. (✿•ᴗ•)\033[0m")
		ui.PrintHorizontalRule()
	} else {
		ui.PrintStatus("analyzing pkgbuild payload signatures carefully...")
		ui.PrintHorizontalRule()
		ui.PrintColorizedPKGBUILD(pkgbuildContent)
		ui.PrintHorizontalRule()

		report := security.ScanPKGBUILD(pkgbuildContent)
		ui.PrintSecurityReport(report)

		if report.RiskLevel != security.Safe {
			fmt.Println("🐾 Security warnings discovered!")
			if ui.AskConfirmation("Do you want to ignore warnings and force step proceed manually?") {
				ui.PrintStatus("bypassing safety trigger guards on user demand...")
			} else {
				ui.PrintCancel()
				return
			}
		}
		_ = security.SaveTrustedHash(pkgHash)
	}

	if ui.AskConfirmation("Wanna Install This AUR as sudo?") {
		ui.PrintSudoWarning()
	}

	ui.PrintStatus("spawning build pipeline compiler subprocess...")
	err = build.BuildAndInstall(dirPath)
	if err != nil {
		ui.PrintFailure()
		return
	}

	if ui.AskConfirmation("Remove unused dependencies after install?") {
		ui.PrintStatus("vacuuming leftover build artifacts...")
		_ = build.CleanupDependencies()
	}

	ui.PrintSuccess()
}
