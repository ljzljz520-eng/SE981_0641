package workflow5

import (
	"context"
	"testing"
	"time"
	"warehouse5s/internal/model"
)

func TestWorkflow5BusinessInvariant(t *testing.T) {
	s := setup(t)
	i := model.NewInspection("1", "A", "I")
	i.AddFinding(model.Finding{Severity: 1})
	s.Record(i)
	ctx, c := context.WithCancel(context.Background())
	c()
	_ = s.RunSteps(ctx, "1")
	time.Sleep(20 * time.Millisecond)
	got, _ := s.Store.Get("1")
	if got.Status != "draft" {
		t.Fatalf("cancelled workflow wrote %s", got.Status)
	}
}
