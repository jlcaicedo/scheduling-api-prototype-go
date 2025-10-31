package schedules

import (
	"testing"
	"time"
)

func TestStore(t *testing.T) {
	st := NewStore()
	if len(st.List()) != 0 {
		t.Fatal("expected empty list")
	}
	s := st.Create("X", time.Unix(0, 0).UTC())
	if s.ID == "" || s.Title != "X" {
		t.Fatal("invalid schedule created")
	}
	if got := len(st.List()); got != 1 {
		t.Fatalf("expected 1, got %d", got)
	}
}
