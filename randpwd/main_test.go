package main

import (
	"testing"
)

func TestRandPassword(t *testing.T) {
	const n = 16
	w1 := randPassword(n)
	if len(w1) != n {
		t.Errorf("%q has %d chars, want 16", w1, n)
	}

	w2 := randPassword(n)
	if w2 == w1 {
		t.Errorf("%q is the same as %q", w1, w2)
	}
}
