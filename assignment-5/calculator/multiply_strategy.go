package calculator

type MultiplyStrategy struct {
}

func (MultiplyStrategy) Execute(a, b float64) (float64, error) {
	return a * b, nil
}
