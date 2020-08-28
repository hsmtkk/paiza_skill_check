package c015_test

import (
	"testing"

	"github.com/hsmtkk/paiza_skill_check/c015"
	"github.com/hsmtkk/paiza_skill_check/test"
)

func Test0(t *testing.T) {
	inputs := []c015.Record{}
	inputs = append(inputs, c015.Record{Date: 1, Amount: 1024})
	inputs = append(inputs, c015.Record{Date: 11, Amount: 2048})
	inputs = append(inputs, c015.Record{Date: 21, Amount: 4196})
	want := 71
	got := c015.New().Calculate(inputs)
	test.AssertEqualInt(t, want, got)
}

func Test1(t *testing.T) {
	inputs := []c015.Record{}
	inputs = append(inputs, c015.Record{Date: 30, Amount: 340})
	want := 10
	got := c015.New().Calculate(inputs)
	test.AssertEqualInt(t, want, got)
}

func Test2(t *testing.T) {
	inputs := []c015.Record{}
	inputs = append(inputs, c015.Record{Date: 5, Amount: 10000})
	inputs = append(inputs, c015.Record{Date: 12, Amount: 10000})
	want := 600
	got := c015.New().Calculate(inputs)
	test.AssertEqualInt(t, want, got)
}
