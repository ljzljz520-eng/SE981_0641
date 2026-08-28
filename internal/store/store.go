package store

import (
	"encoding/json"
	"go.etcd.io/bbolt"
	"os"
	"sync"
	"warehouse5s/internal/model"
)

var bucket = []byte("inspections")

type Store struct {
	db *bbolt.DB
	mu sync.Mutex
}

func Open(path string) (*Store, error) {
	db, e := bbolt.Open(path, 0600, nil)
	if e != nil {
		return nil, e
	}
	s := &Store{db: db}
	e = db.Update(func(tx *bbolt.Tx) error { _, x := tx.CreateBucketIfNotExists(bucket); return x })
	if e != nil {
		db.Close()
		return nil, e
	}
	return s, nil
}
func (s *Store) Close() error { return s.db.Close() }
func (s *Store) Save(i model.Inspection) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	b, e := json.Marshal(i)
	if e != nil {
		return e
	}
	return s.db.Update(func(tx *bbolt.Tx) error { return tx.Bucket(bucket).Put([]byte(i.ID), b) })
}
func (s *Store) Get(id string) (model.Inspection, error) {
	var i model.Inspection
	e := s.db.View(func(tx *bbolt.Tx) error {
		v := tx.Bucket(bucket).Get([]byte(id))
		if v == nil {
			return os.ErrNotExist
		}
		return json.Unmarshal(v, &i)
	})
	return i, e
}
func (s *Store) List() ([]model.Inspection, error) {
	out := []model.Inspection{}
	e := s.db.View(func(tx *bbolt.Tx) error {
		return tx.Bucket(bucket).ForEach(func(_, v []byte) error {
			var i model.Inspection
			if e := json.Unmarshal(v, &i); e != nil {
				return e
			}
			out = append(out, i)
			return nil
		})
	})
	return out, e
}
func (s *Store) Delete(id string) error {
	return s.db.Update(func(tx *bbolt.Tx) error { return tx.Bucket(bucket).Delete([]byte(id)) })
}
