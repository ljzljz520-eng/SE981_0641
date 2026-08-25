package policy

import (
	"testing"
	"warehouse5s/internal/model"
)

func TestPolicyValidation(t *testing.T) {
	i := model.NewInspection("1", "A", "I")
	if len(Validate(i)) == 0 {
		t.Fatal()
	}
}
