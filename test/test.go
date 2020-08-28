package test

import "testing"

func AssertEqualInt(t *testing.T, want, got int) {
	t.Helper()
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}
