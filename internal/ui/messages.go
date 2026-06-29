package ui

import (
	"fmt"
	"strings"
	"ncah/internal/aur"
	"ncah/internal/security"
)

const (
	Cyan        = "\033[0;36m"
	Green       = "\033[0;32m"
	Yellow      = "\033[0;33m"
	Red         = "\033[0;31m"
	DarkGray    = "\033[1;30m"
	White       = "\033[0;37m"
	BoldWhite   = "\033[1;37m"
	Reset       = "\033[0m"
)

func PrintStatus(msg string) {
	fmt.Printf("%sNCAH 🐾%s %s\n", Green, Reset, msg)
}

func PrintHorizontalRule() {
	fmt.Printf("%s──────────────────────────────────────────────────────────────%s\n", DarkGray, Reset)
}

func PrintFailure() {
	fmt.Printf("\n%sGomennasai~~ The build system ran into a big oofie... ｡°(°.◜ᯅ◝°)°｡%s\n", Red, Reset)
}

func PrintCancel() {
	fmt.Printf("\n%sEhhh... Are You Cancel This Transaction Nyaww~~ (｡>﹏<｡)%s\n", Red, Reset)
}

func PrintSuccess() {
	fmt.Printf("\n%sYatta~! The package was successfully installed nyaa~! Besties forever! 🎉✨ (っ^‿^)っ%s\n", Green, Reset)
}

func PrintSudoWarning() {
	fmt.Printf("%s[!] \"B-but master... NCAH needs your superuser password pwd to continue nyaa~!\"%s\n", Yellow, Reset)
}

func PrintSearchResults(results []aur.AURPackage) {
	if len(results) == 0 {
		fmt.Printf("%sGomennasai~ Couldn't find any cute packages matching that (っ- ‸ -,)%s\n", Yellow, Reset)
		return
	}
	for _, pkg := range results {
		fmt.Printf("%saur/%s%s %s%s%s\n    💌 %s%s\n", Cyan, BoldWhite, pkg.Name, Reset, Green, pkg.Version, Reset, pkg.Description)
	}
}

func PrintPackageInfo(pkg *aur.AURPackage) {
	PrintHorizontalRule()
	fmt.Printf("%s📂 Super Detailed Metadata Profile for Bestie: %s%s\n", Cyan, BoldWhite, pkg.Name)
	PrintHorizontalRule()
	fmt.Printf("%s  • Package Base    :%s %s\n", Cyan, Reset, pkg.PackageBase)
	fmt.Printf("%s  • Adoption Version:%s %s\n", Cyan, Reset, pkg.Version)
	fmt.Printf("%s  • Description     :%s %s\n", Cyan, Reset, pkg.Description)
	fmt.Printf("%s  • Upstream Link   :%s %s\n", Cyan, Reset, pkg.URL)
	fmt.Printf("%s  • Source Git Repo :%s https://aur.archlinux.org/%s.git\n", Cyan, Reset, pkg.PackageBase)
	fmt.Printf("%s  • Head Pat Giver  :%s %s\n", Cyan, Reset, pkg.Maintainer)
	fmt.Printf("%s  • License Profile :%s %s\n", Cyan, Reset, stringify(pkg.License))
	fmt.Printf("%s  • Popularity Score:%s %.2f 🔥\n", Cyan, Reset, pkg.Popularity)
	fmt.Printf("%s  • Community Votes :%s %d 🗳️\n", Cyan, Reset, pkg.NumVotes)

	if len(pkg.Depends) > 0 {
		fmt.Printf("%s  • Core Buddies    :%s %s\n", Cyan, Reset, stringify(pkg.Depends))
	}
	if len(pkg.MakeDepends) > 0 {
		fmt.Printf("%s  • Build Buddies   :%s %s\n", Cyan, Reset, stringify(pkg.MakeDepends))
	}
	if len(pkg.OptDepends) > 0 {
		fmt.Printf("%s  • Optional Extras :%s %s\n", Cyan, Reset, stringify(pkg.OptDepends))
	}
	PrintHorizontalRule()
}

func PrintDependencies(depends []string, makeDepends []string) {
	if len(depends) > 0 {
		fmt.Printf("  📦 Core Buddies: %s%s%s\n", Cyan, stringify(depends), Reset)
	}
	if len(makeDepends) > 0 {
		fmt.Printf("  🛠️  Build Buddies: %s%s%s\n", Yellow, stringify(makeDepends), Reset)
	}
	if len(depends) == 0 && len(makeDepends) == 0 {
		fmt.Printf("  %sWoww! This package is super independent, no besties needed meww~%s\n", Green, Reset)
	}
}

func PrintColorizedPKGBUILD(content string) {
	lines := strings.Split(content, "\n")
	for i, line := range lines {
		lineNo := i + 1
		trimmed := strings.TrimSpace(line)
		fmt.Printf("%s%3d │ %s", DarkGray, lineNo, Reset)

		if strings.HasPrefix(trimmed, "#") {
			fmt.Printf("%s%s%s\n", DarkGray, line, Reset)
		} else if strings.Contains(line, "=") && !strings.Contains(line, "()") {
			parts := strings.SplitN(line, "=", 2)
			fmt.Printf("%s%s%s=%s%s%s\n", Cyan, parts[0], Reset, White, parts[1], Reset)
		} else if strings.Contains(line, "()") || strings.HasPrefix(trimmed, "function ") {
			fmt.Printf("%s%s%s\n", Green, line, Reset)
		} else {
			fmt.Printf("%s%s%s\n", White, line, Reset)
		}
	}
}

func PrintSecurityReport(report *security.ScanReport) {
	var color string
	switch report.RiskLevel {
	case security.Safe:
		color = fmt.Sprintf("%s[SAFE] ✨ NCAH core loves this wholesome code signature!%s", Green, Reset)
	case security.Warning:
		color = fmt.Sprintf("%s[WARNING] ⚠️ Mindful inspection required, nyaa!%s", Yellow, Reset)
	case security.HighRisk:
		color = fmt.Sprintf("%s[HIGH RISK] 🔥🙀 Dangerous anomalies detected rawr!%s", Red, Reset)
	}

	fmt.Printf("\n🛡️  %sNCAH Security Assessment Tools:%s %s\n", BoldWhite, Reset, color)
	if len(report.Findings) > 0 {
		for _, finding := range report.Findings {
			var rColor string
			if finding.RiskLevel == security.HighRisk {
				rColor = fmt.Sprintf("%sHIGH%s", Red, Reset)
			} else {
				rColor = fmt.Sprintf("%sWARNING%s", Yellow, Reset)
			}
			fmt.Printf("%s[!] Risk: %s%s\n", Red, rColor, Reset)
			fmt.Printf("    %sline %d:%s %s%s%s\n", DarkGray, finding.LineNo, Reset, White, finding.MatchedText, Reset)
			fmt.Printf("    %sreason:%s %s\n\n", Cyan, Reset, finding.Reason)
		}
	} else {
		fmt.Printf("    %s✔ Squeaky clean! Wholesome code signature verified nyaa~%s\n", Green, Reset)
	}
}

func stringify(slice []string) string {
	if len(slice) == 0 {
		return "None meww"
	}
	return "[" + strings.Join(slice, ", ") + "]"
}
