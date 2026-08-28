package archive

import (
	"warehouse5s/internal/model"
	"warehouse5s/internal/store"
)

func Eligible(i model.Inspection) bool { return i.Status == "approved" }
func Mark(s *store.Store, i model.Inspection, reason string) error {
	if !Eligible(i) {
		return nil
	}
	i.Status = "archived"
	return s.Save(i)
}
func RetentionClass(reason string) string {
	if reason == "legal" {
		return "permanent"
	}
	return "standard"
}
