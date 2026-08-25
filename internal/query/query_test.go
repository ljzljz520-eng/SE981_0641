package query

import (
	"testing"
	"warehouse5s/internal/model"
	"warehouse5s/internal/store"
)

func TestSearchByArea(t *testing.T) {
	s, _ := store.Open(t.TempDir() + "/x")
	defer s.Close()
	s.Save(model.NewInspection("1", "A", "I"))
	x, e := Search(s, "A", "")
	if e != nil || len(x) != 1 {
		t.Fatal()
	}
}
