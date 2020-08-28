package c015

import (
	"strconv"
	"strings"
)

type Record struct {
	Date   int
	Amount int
}

type Calculator interface {
	Calculate([]Record) int
}

type calculatorImpl struct{}

func New() Calculator {
	return &calculatorImpl{}
}

func (c *calculatorImpl) Calculate(records []Record) int {
	points := 0
	for _, r := range records {
		if IncludeThree(r.Date) {
			points += int(r.Amount * 3 / 100)
		} else if IncludeFive(r.Date) {
			points += int(r.Amount * 5 / 100)
		} else {
			points += int(r.Amount / 100)
		}
	}
	return points
}

func IncludeThree(n int) bool {
	s := strconv.Itoa(n)
	return strings.Contains(s, "3")
}

func IncludeFive(n int) bool {
	s := strconv.Itoa(n)
	return strings.Contains(s, "5")
}
