package c023_test

import (
	"testing"

	"github.com/hsmtkk/paiza_skill_check/c023"
	"github.com/hsmtkk/paiza_skill_check/test"
)

func Test0(t *testing.T) {
	counter := c023.New([]int{1, 2, 3, 4, 5, 6})

	want := 6
	got := counter.Count([]int{1, 5, 4, 2, 3, 6})
	test.AssertEqualInt(t, want, got)

	want = 4
	got = counter.Count([]int{9, 6, 2, 7, 1, 5})
	test.AssertEqualInt(t, want, got)

	want = 0
	got = counter.Count([]int{32, 9, 87, 33, 41, 60})
	test.AssertEqualInt(t, want, got)
}
