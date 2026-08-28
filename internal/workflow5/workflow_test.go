package workflow5

import (
	"os"
	"testing"
	"warehouse5s/internal/model"
	"warehouse5s/internal/store"
)

func setup(t *testing.T) *Service {
	p := t.TempDir() + "/x.db"
	s, e := store.Open(p)
	if e != nil {
		t.Fatal(e)
	}
	t.Cleanup(func() { s.Close(); os.Remove(p) })
	return New(s)
}
func TestWorkflowOne(t *testing.T) {
	s := setup(t)
	i := model.NewInspection("1", "A", "I")
	i.AddFinding(model.Finding{Severity: 1})
	if s.Record(i) != nil || s.Submit("1") != nil {
		t.Fatal()
	}
}
func TestWorkflowTwo(t *testing.T) {
	s := setup(t)
	i := model.NewInspection("1", "A", "I")
	i.AddFinding(model.Finding{Severity: 1})
	s.Record(i)
	s.Submit("1")
	if s.Review("1", "Q", true) != nil {
		t.Fatal()
	}
}
func TestWorkflowThree(t *testing.T) {
	s := setup(t)
	i := model.NewInspection("1", "A", "I")
	i.AddFinding(model.Finding{Severity: 1})
	s.Record(i)
	s.Submit("1")
	s.Review("1", "Q", true)
	if s.Archive("1", "r") != nil {
		t.Fatal()
	}
}
