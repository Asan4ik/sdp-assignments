package calculator

import "errors"

var ErrNoStrategy = errors.New("calculator: no strategy set")

type Calculator struct {
	strategy OperationStrategy
}

func NewCalculator(s OperationStrategy) *Calculator {
	return &Calculator{strategy: s}
}

func (c *Calculator) SetStrategy(s OperationStrategy) {
	c.strategy = s
}

func (c *Calculator) Calculate(a, b float64) (float64, error) {
	if c.strategy == nil {
		return 0, ErrNoStrategy
	}
	return c.strategy.Execute(a, b)
}
