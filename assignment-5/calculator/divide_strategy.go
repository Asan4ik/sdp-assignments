package calculator

import "errors"

var ErrDivisionByZero = errors.New("division by zero")

type DivideStrategy struct{}

func (DivideStrategy) Execute(a, b float64) (float64, error) {
	if b == 0 {
		return 0, ErrDivisionByZero
	}
	return a / b, nil
}
