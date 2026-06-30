package main

import (
	"fmt"
	"os"
	"strings"

	"ncah/internal/aur"
	"ncah/internal/build"
	"ncah/internal/ui"
)

func main() {
	ui.PrintBanner()

	if len(os.Args) < 2 {
		fmt.Println("👉 pssssst! use -Ss [query] to hunt, -Si [pkg] to peek, -S [pkg] to adopt, or -R [pkg] to banish a package nyaa~! (✿•ᴗ•)")
		return
	}

	arg := os.Args[1]

	switch arg {
	case "-Sy", "-Syy", "-Su", "-Syu", "-Syyu":
		handleSystemSync(arg)
		return
	case "-Ss":
		if len(os.Args) < 3 {
			fmt.Println("❌ Missing search query nyaa~!")
			return
		}
		handleSearch(os.Args[2])
		return
	case "-Si":
		if len(os.Args) < 3 {
			fmt.Println("❌ Missing package name nyaa~!")
			return
		}
		handleInfo(os.Args[2])
		return
	case "-R":
		if len(os.Args) < 3 {
			fmt.Println("❌ Missing package name nyaa~!")
			return
		}
		handleRemove(os.Args[2])
		return
	case "-S":
		if len(os.Args) < 3 {
			fmt.Println("❌ Missing package name nyaa~!")
			return
		}
		handleInstall(os.Args[2])
		return
	default:
		if strings.HasPrefix(arg, "-S") {
			handleSystemSync(arg)
			return
		}
		fmt.Println("👉 pssssst! use -Ss [query] to hunt, -Si [pkg] to peek, -S [pkg] to adopt, or -R [pkg] to banish a package nyaa~! (✿•ᴗ•)")
	}
}

func handleSystemSync(flag string) {
	ui.PrintStatus(fmt.Sprintf("initiating official repository synchronization using %s protocol nyaa~ ✨", flag))
	ui.PrintSudoWarning()
	err := build.SyncOfficial(flag)
	if err != nil {
		ui.PrintCancel()
		return
	}

	if strings.Contains(flag, "u") {
		ui.PrintHorizontalRule()
		ui.PrintStatus("checking for available AUR package upgrades... 🐾")

		foreign, err := build.GetForeignPackages()
		if err != nil {
			ui.PrintFailure()
			return
		}

		var upgradable []string
		for name, localVer := range foreign {
			info, err := aur.GetPackageInfo(name)
			if err != nil || info == nil {
				continue
			}
			if build.CheckVersionUpgrade(localVer, info.Version) {
				fmt.Printf("  \033[0;33m⚡ Upgrade found:\033[0m %s [%s -> %s]\n", name, localVer, info.Version)
				upgradable = append(upgradable, name)
			}
		}

		if len(upgradable) == 0 {
			fmt.Println("  \033[0;32m✔ All AUR packages are perfectly up to date and cozy master! ˚ʚ♡ɞ˚\033[0m")
			ui.PrintSuccess()
			return
		}

		ui.PrintHorizontalRule()
		if !ui.AskConfirmation(fmt.Sprintf("Wanna deploy upgrade protocols for all %d out-of-date AUR packages meww~?", len(upgradable))) {
			ui.PrintCancel()
			return
		}

		for _, name := range upgradable {
			ui.PrintHorizontalRule()
			ui.PrintStatus(fmt.Sprintf("processing obsessive upgrade for %s...", name))
			handleInstall(name)
		}
	}
	ui.PrintSuccess()
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

	ui.PrintStatus("analyzing pkgbuild payload signatures carefully... 🕵️‍♀️")
	ui.PrintHorizontalRule()

	findings := build.ScanMalware(pkgbuildContent)
	if len(findings) > 0 {
		fmt.Println("\n\033[0;31m🔥🙀 [HIGH RISK] MALWARE INTRUSION ALERTS DETECTED RAWR!:\033[0m")
		for _, line := range findings {
			fmt.Printf("  \033[0;31m⚡ Warning match found:\033[0m %s\n", line)
		}
		ui.PrintHorizontalRule()
		if !ui.AskConfirmation("Dangerous anomalies detected! Do you still want to bypass safety guards and proceed?") {
			ui.PrintCancel()
			return
		}
	} else {
		fmt.Println("  \033[0;32m✔ Wholesome code signature verified nyaa~! Pure love! ˚ʚ♡ɞ˚\033[0m")
		ui.PrintHorizontalRule()
	}

	ui.PrintStatus("deploying ClamAV signature database shield... 🛡️")
	virusLog, virusFound := build.ScanVirus(dirPath)
	if virusFound {
		fmt.Println("\n\033[0;31m🔥🙀 This AUR Was Infected nyaww~~~\033[0m")
		fmt.Println(virusLog)
		ui.PrintHorizontalRule()
		ui.PrintCancel()
		return
	} else if virusLog != "" {
		fmt.Printf("  \033[0;33m⚠️ %s\033[0m\n", virusLog)
		ui.PrintHorizontalRule()
	} else {
		fmt.Println("  \033[0;32m✔ No One Virus Are Detected nyaa~~~\033[0m")
		ui.PrintHorizontalRule()
	}

	if ui.AskConfirmation("Wanna inspect or edit the PKGBUILD before compiling nyaa~?") {
		err = build.EditPKGBUILD(dirPath)
		if err != nil {
			ui.PrintFailure()
			return
		}
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