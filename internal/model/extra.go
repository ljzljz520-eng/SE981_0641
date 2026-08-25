package model

func (f Finding) IsCritical() bool     { return f.Severity >= 4 }
func (i Inspection) FindingCount() int { return len(i.Findings) }
func (i *Inspection) Resolve(id string) bool {
	for n := range i.Findings {
		if i.Findings[n].ID == id {
			i.Findings[n].Resolved = true
			return true
		}
	}
	return false
}
