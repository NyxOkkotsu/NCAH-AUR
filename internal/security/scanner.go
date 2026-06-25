package security

type ScanReport struct {
	RiskLevel RiskLevel
	Findings  []string
}

func ScanPKGBUILD(content string) *ScanReport {
	report := &ScanReport{
		RiskLevel: Safe,
		Findings:  []string{},
	}

	rules := GetSecurityRules()

	for _, rule := range rules {
		if rule.Pattern.MatchString(content) {
			report.Findings = append(report.Findings, rule.Description)
			if rule.Risk > report.RiskLevel {
				report.RiskLevel = rule.Risk
			}
		}
	}

	return report
}