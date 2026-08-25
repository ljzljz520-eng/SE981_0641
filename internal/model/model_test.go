package model

import "testing"

func TestInspectionScoring(t *testing.T) {
	i := NewInspection("1", "A", "x")
	i.AddFinding(Finding{Severity: 2})
	if i.CalculateScore() != 80 {
		t.Fatal()
	}
}
