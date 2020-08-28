package c076

type TimeSalary struct {
	Day      int
	Night    int
	MidNight int
}

type WorkingHour struct {
	Begin int
	End   int
}

type Calculator interface {
	Calculate(TimeSalary, []WorkingHour) int
}

type calcImpl struct{}

func New() Calculator {
	return &calcImpl{}
}

func (c *calcImpl) Calculate(ts TimeSalary, hours []WorkingHour) int {
	sum := 0
	for _, hour := range hours {
		// midnight
		for i := 0; i < 9; i++ {
			if hour.Begin <= i && i < hour.End {
				sum += ts.MidNight
			}
		}
		// day
		for i := 9; i < 17; i++ {
			if hour.Begin <= i && i < hour.End {
				sum += ts.Day
			}
		}
		// night
		for i := 17; i < 22; i++ {
			if hour.Begin <= i && i < hour.End {
				sum += ts.Night
			}
		}
		// midnight
		for i := 22; i < 24; i++ {
			if hour.Begin <= i && i < hour.End {
				sum += ts.MidNight
			}
		}
	}
	return sum
}
