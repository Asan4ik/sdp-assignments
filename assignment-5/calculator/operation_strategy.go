package calculator

// OperationStrategy is the Strategy interface. Every arithmetic algorithm
// implements Execute, so the Calculator can use them interchangeably.
type OperationStrategy interface {
	Execute(a, b float64) (float64, error)
}

var (
	_ OperationStrategy = AddStrategy{}
	_ OperationStrategy = SubtractStrategy{}
	_ OperationStrategy = MultiplyStrategy{}
	_ OperationStrategy = DivideStrategy{}
)
