package calculator

type AddStrategy struct {
}

func (AddStrategy) Execute(a, b float64) float64 {
	return a + b
}
