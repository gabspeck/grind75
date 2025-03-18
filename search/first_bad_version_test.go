package search

import "testing"

func TestFirstBadVersionAttempt(t *testing.T) {
	result := FirstBadVersionAttempt(50_000)
	if result != 4 {
		t.Errorf("wanted %v, got %v", 4, result)
	}
}
