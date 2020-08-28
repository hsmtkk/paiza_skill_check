package c023

type Counter interface {
	Count([]int) int
}

type counterImpl struct {
	correctNumbers []int
}

func New(correctNumbers []int) Counter {
	return &counterImpl{correctNumbers: correctNumbers}
}

func (c *counterImpl) Count(nums []int) int {
	count := 0
	for _, cn := range c.correctNumbers {
		for _, n := range nums {
			if cn == n {
				count++
			}
		}
	}
	return count
}
