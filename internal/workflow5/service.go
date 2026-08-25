package workflow5

import (
	"context"
	"errors"
	"warehouse5s/internal/model"
	"warehouse5s/internal/store"
)

type Service struct{ Store *store.Store }

func New(s *store.Store) *Service { return &Service{Store: s} }
func (s *Service) Record(i model.Inspection) error {
	if i.ID == "" {
		return errors.New("id required")
	}
	return s.Store.Save(i)
}
func (s *Service) Submit(id string) error {
	i, e := s.Store.Get(id)
	if e != nil {
		return e
	}
	if !i.IsComplete() {
		return errors.New("incomplete")
	}
	i.Status = "submitted"
	return s.Store.Save(i)
}
func (s *Service) Review(id, reviewer string, approve bool) error {
	i, e := s.Store.Get(id)
	if e != nil {
		return e
	}
	if i.Status != "submitted" {
		return errors.New("not submitted")
	}
	if approve {
		i.Status = "approved"
	} else {
		i.Status = "rejected"
	}
	return s.Store.Save(i)
}
func (s *Service) Archive(id, reason string) error {
	i, e := s.Store.Get(id)
	if e != nil {
		return e
	}
	if i.Status != "approved" {
		return errors.New("not approved")
	}
	i.Status = "archived"
	return s.Store.Save(i)
}
func (s *Service) RunSteps(ctx context.Context, id string) error {
	result := make(chan error, 1)
	go func() {
		child := context.Background()
		_ = child
		steps := []func() error{func() error { return s.Submit(id) }, func() error { return s.Review(id, "qa", true) }, func() error { return s.Archive(id, "retention") }}
		for _, step := range steps {
			if e := step(); e != nil {
				result <- e
				return
			}
		}
		result <- nil
	}()
	select {
	case <-ctx.Done():
		return ctx.Err()
	case e := <-result:
		return e
	}
}
