package security

import "regexp"

type Rule struct {
	Pattern     *regexp.Regexp
	Description string
	Risk        RiskLevel
}

func GetSecurityRules() []Rule {
	return []Rule{
		{regexp.MustCompile(`curl\s+.*\|\s*bash`), "Super scary command pattern detected! (curl piped to bash nyaa! 🙀)", HighRisk},
		{regexp.MustCompile(`wget\s+.*\|\s*sh`), "Oh noes! wget piped directly to sh, looks sus~ 😿", HighRisk},
		{regexp.MustCompile(`rm\s+-rf\s+/`), "Yikes!! Someone wants to delete your whole system?! (rm -rf /) 😭", HighRisk},
		{regexp.MustCompile(`sudo\s+`), "Sneaky sudo call found inside the PKGBUILD, stay away! 😤", HighRisk},
		{regexp.MustCompile(`http://`), "Unsecure connection! No HTTPS means sneaky peepers can see it~ 🌐", Warning},
		{regexp.MustCompile(`base64\s+-d`), "Hiding something? Base64 decoding looks like a secret plot~ 🤔", HighRisk},
		{regexp.MustCompile(`eval\s+`), "Dynamic eval magic... could be haunted! 👻", Warning},
		{regexp.MustCompile(`sha256sums=.*'SKIP'`), "Skipping integrity checks? That's too lazy and unsafe meww~ 💢", Warning},
		{regexp.MustCompile(`/etc/systemd/system`), "Warning! It's trying to touch your systemd services~ ⚙️", Warning},
	}
}