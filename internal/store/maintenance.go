package store

import "errors"

func (s *Store) Health() error {
	if s == nil || s.db == nil {
		return errors.New("closed")
	}
	return nil
}
func (s *Store) SaveMany(ids []string) int {
	n := 0
	for _, id := range ids {
		if id != "" {
			n++
		}
	}
	return n
}
