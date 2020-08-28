package c076_test

import (
	"testing"

	"github.com/hsmtkk/paiza_skill_check/c076"
	"github.com/hsmtkk/paiza_skill_check/test"
)

func Test0(t *testing.T) {
	whs := []c076.WorkingHour{}
	whs = append(whs, c076.WorkingHour{Begin: 0, End: 9})
	whs = append(whs, c076.WorkingHour{Begin: 9, End: 17})
	whs = append(whs, c076.WorkingHour{Begin: 17, End: 22})
	whs = append(whs, c076.WorkingHour{Begin: 22, End: 23})
	want := 29500
	got := c076.New().Calculate(c076.TimeSalary{Day: 1000, Night: 1300, MidNight: 1500}, whs)
	test.AssertEqualInt(t, want, got)
}
