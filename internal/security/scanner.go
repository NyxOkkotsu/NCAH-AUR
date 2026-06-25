package security

import "strings"

type Finding struct {
	LineNo      int
	MatchedText string
	Reason      string
	RiskLevel   RiskLevel
}

type ScanReport struct {
	RiskLevel RiskLevel
	Findings  []Finding
}

func ScanPKGBUILD(content string) *ScanReport {
	report := &ScanReport{
		RiskLevel: Safe,
		Findings:  []Finding{},
	}

	rules := GetSecurityRules()
	lines := strings.Split(content, "\n")

	for lineIdx, line := range lines {
		trimmedLine := strings.TrimSpace(line)
		if trimmedLine == "" || strings.HasPrefix(trimmedLine, "#") {
			continue
		}

		for _, rule := range rules {
			if rule.Pattern.MatchString(trimmedLine) {
				finding := Finding{
					LineNo:      lineIdx + 1,
					MatchedText: trimmedLine,
					Reason:      rule.Reason,
					RiskLevel:   rule.Risk,
				}
				report.Findings = append(report.Findings, finding)
				if rule.Risk > report.RiskLevel {
					report.RiskLevel = rule.Risk
				}
			}
		}
	}
	return report
}