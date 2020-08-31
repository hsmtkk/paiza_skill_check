package b079

type Calculator interface {
	Calculate(string, string) int
}

type calcImpl struct{}

func New() Calculator {
	return &calcImpl{}
}

func (c *calcImpl) Calculate(name0, name1 string) int {
	score0 := calculate(name0 + name1)
	score1 := calculate(name1 + name0)
	if score0 > score1 {
		return score0
	}
	return score1
}

func calculate(s string) int {
	ns := []int{}
	for _, ch := range s {
		ns = append(ns, ch2int(ch))
	}
	for {
		if len(ns) == 1 {
			break
		}
		ns = calculate2(ns)
	}
	return ns[0]
}

func ch2int(r rune) int {
	r0 := 'a'
	return int(r-r0) + 1
}

func calculate2(ns []int) []int {
	ret := []int{}
	for i := 0; i < len(ns)-1; i++ {
		sum := ns[i] + ns[i+1]
		if sum > 101 {
			sum -= 101
		}
		ret = append(ret, sum)
	}
	return ret
}
