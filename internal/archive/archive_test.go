package archive

import (
	"testing"
	"warehouse5s/internal/model"
)

func TestRetentionClass(t *testing.T) {
	if RetentionClass("legal") != "permanent" {
		t.Fatal()
	}
	if !Eligible(model.Inspection{Status: "approved"}) {
		t.Fatal()
	}
}
