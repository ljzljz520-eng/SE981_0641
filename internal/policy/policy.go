package policy

import "warehouse5s/internal/model"

func Validate(i model.Inspection) []string {
	issues := []string{}
	if i.Area == "" {
		issues = append(issues, "area")
	}
	if len(i.Findings) == 0 {
		issues = append(issues, "findings")
	}
	for _, f := range i.Findings {
		if f.Severity < 1 || f.Severity > 5 {
			issues = append(issues, "severity")
		}
	}
	return issues
}
func CanApprove(i model.Inspection) bool { return i.Status == "submitted" && len(Validate(i)) == 0 }
func RequiresEscalation(i model.Inspection) bool {
	for _, f := range i.Findings {
		if f.Severity >= 4 {
			return true
		}
	}
	return false
}
