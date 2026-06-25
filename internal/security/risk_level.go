package security

type RiskLevel int

const (
	Safe RiskLevel = iota
	Warning
	HighRisk
)