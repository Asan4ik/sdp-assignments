package calculator

type SubtractStrategy struct {
}

func (SubtractStrategy) Execute(a, b float64) float64 {
	return a - b
}
