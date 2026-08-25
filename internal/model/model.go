package model

type Inspection struct {
	ID, Area, Inspector string
	Status              string
	Score               int
	Findings            []Finding
}
type Finding struct {
	ID, InspectionID, Description, Category string
	Severity                                int
	Resolved                                bool
}
type Approval struct{ ID, InspectionID, Reviewer, Decision string }
type ArchiveRecord struct{ ID, InspectionID, Reason string }
type EntityOne struct{ ID string }
type EntityTwo struct{ ID string }
type EntityThree struct{ ID string }
type EntityFour struct{ ID string }

func NewInspection(id, area, inspector string) Inspection {
	return Inspection{ID: id, Area: area, Inspector: inspector, Status: "draft"}
}
func (i Inspection) IsComplete() bool {
	return i.ID != "" && i.Area != "" && i.Inspector != "" && len(i.Findings) > 0
}
func (i *Inspection) AddFinding(f Finding) {
	if f.ID == "" {
		f.ID = i.ID + "-" + string(rune(len(i.Findings)+48))
	}
	f.InspectionID = i.ID
	i.Findings = append(i.Findings, f)
}
func (i Inspection) CalculateScore() int {
	score := 100
	for _, f := range i.Findings {
		score -= f.Severity * 10
	}
	if score < 0 {
		return 0
	}
	return score
}
