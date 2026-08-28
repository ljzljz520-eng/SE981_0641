package query

import (
	"strings"
	"warehouse5s/internal/model"
	"warehouse5s/internal/store"
)

func Search(s *store.Store, area, status string) ([]model.Inspection, error) {
	all, e := s.List()
	if e != nil {
		return nil, e
	}
	out := []model.Inspection{}
	for _, i := range all {
		if area != "" && !strings.Contains(i.Area, area) {
			continue
		}
		if status != "" && i.Status != status {
			continue
		}
		out = append(out, i)
	}
	return out, nil
}
func ByInspector(s *store.Store, name string) ([]model.Inspection, error) { return Search(s, name, "") }
func CountByStatus(items []model.Inspection, status string) int {
	n := 0
	for _, i := range items {
		if i.Status == status {
			n++
		}
	}
	return n
}
