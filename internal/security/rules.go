package security

import "regexp"

type RiskLevel int

const (
	Safe RiskLevel = iota
	Warning
	HighRisk
)

type Rule struct {
	Pattern *regexp.Regexp
	Reason  string
	Risk    RiskLevel
}

func GetSecurityRules() []Rule {
	return []Rule{
		{regexp.MustCompile(`curl\s+.*\|\s*bash`), "remote script execution via piping curl to bash", HighRisk},
		{regexp.MustCompile(`wget\s+.*\|\s*sh`), "remote script execution via piping wget to sh", HighRisk},
		{regexp.MustCompile(`rm\s+-rf\s+/`), "malicious root directory deletion attempt", HighRisk},
		{regexp.MustCompile(`sudo\s+`), "unsafe manual escalation boundary command inside packaging rules", HighRisk},
		{regexp.MustCompile(`http://`), "unencrypted remote resource retrieval source link used", Warning},
		{regexp.MustCompile(`base64\s+-d`), "dangerous obfuscation payload pattern mechanism detected", HighRisk},
		{regexp.MustCompile(`eval\s+`), "dynamic evaluation execution backend engine loaded", Warning},
		{regexp.MustCompile(`sha256sums=.*'SKIP'`), "bypassed check payload integrity verification source hashes", Warning},
		{regexp.MustCompile(`/etc/systemd/system`), "sensitive configuration zone path modification manipulation", Warning},
	}
}