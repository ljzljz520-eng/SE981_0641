package store

func (s *Store) SnapshotSize() int {
	xs, e := s.List()
	if e != nil {
		return 0
	}
	return len(xs)
}
