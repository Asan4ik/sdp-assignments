package calculator

type AddStrategy struct {
}

func (AddStrategy) Execute(a, b float64) (float64, error) {
	return a + b, nil
}
