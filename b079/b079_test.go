package b079_test

import (
	"testing"

	"github.com/hsmtkk/paiza_skill_check/b079"
	"github.com/hsmtkk/paiza_skill_check/test"
)

func Test0(t *testing.T) {
	want := 83
	got := b079.New().Calculate("pa", "iza")
	test.AssertEqualInt(t, want, got)
}

func Test1(t *testing.T) {
	want := 97
	got := b079.New().Calculate("alice", "bob")
	test.AssertEqualInt(t, want, got)
}
