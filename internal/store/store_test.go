package store

import (
	"os"
	"testing"
	"warehouse5s/internal/model"
)

func TestPersistenceSurvivesReopen(t *testing.T) {
	p := "test.db"
	defer os.Remove(p)
	s, _ := Open(p)
	s.Save(model.NewInspection("x", "A", "I"))
	s.Close()
	s, _ = Open(p)
	defer s.Close()
	if _, e := s.Get("x"); e != nil {
		t.Fatal(e)
	}
}
