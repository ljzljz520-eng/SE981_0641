package audit

import "testing"

func TestAuditRecords(t *testing.T) {
	l := New()
	l.Record("x", "y", "z")
	if len(l.Events()) != 1 {
		t.Fatal()
	}
}
