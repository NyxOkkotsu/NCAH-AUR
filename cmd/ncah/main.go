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
		ui.PrintStatus("pouncing on the database to sniff out matching packages rawr... 🐾")
		handleSearch(*searchFlag)
		return
	}

	if *infoFlag != "" {
		ui.PrintStatus("peeking at the super secret package metadata definitions uwu... 📂")
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

func handleRemove(pkgName string) {
	ui.PrintStatus(fmt.Sprintf("preparing uninstallation protocol to banish %s nyaa...", pkgName))
	if !ui.AskConfirmation(fmt.Sprintf("Are you absolutely sure you want to purge %s and its configuration remnants from our cozy layout?", pkgName)) {
		fmt.Println("Action aborted! package remains cozy and untouched meww~ (ς°´o`°ς)")
		return
	}
	err := build.RemovePackage(pkgName)
	if err != nil {
		ui.PrintFailure()
		return
	}
	fmt.Printf("\n\033[0;32m✔ %s has been successfully vaporized from the system layout! nyaa~\033[0m\n", pkgName)
}

func handleInstall(pkgName string) {
	ui.PrintStatus("waking up remote AUR api mirrors... mwah")
	info, err := aur.GetPackageInfo(pkgName)
	if err != nil || info == nil {
		ui.PrintFailure()
		return
	}

	fmt.Printf("\n✨ \033[1;36m[Dependencies validation target for: %s]\033[0m\n", pkgName)
	ui.PrintDependencies(info.Depends, info.MakeDepends)

	if !ui.AskConfirmation("Proceed with downloading installation recipes?") {
		fmt.Println("Gumuuu... Action aborted by user! (っ°´o`°ς)")
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
				fmt.Println("Phew~ Smart choice! Application deployment abandoned safely! 🛡️( •̀ ω •́ )✧")
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