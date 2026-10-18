package calculator

type OperationStrategy interface {
	Execute(a, b float64) float64
}
