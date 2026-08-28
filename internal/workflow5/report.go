package workflow5

import (
	"warehouse5s/internal/model"
	"warehouse5s/internal/query"
	"warehouse5s/internal/store"
)

func Report(s *store.Store, area string) (int, error) {
	xs, e := query.Search(s, area, "")
	if e != nil {
		return 0, e
	}
	total := 0
	for _, i := range xs {
		total += i.CalculateScore()
	}
	if len(xs) == 0 {
		return 0, nil
	}
	return total / len(xs), nil
}
func Normalize(i model.Inspection) model.Inspection {
	if i.Score == 0 {
		i.Score = i.CalculateScore()
	}
	return i
}
