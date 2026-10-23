package main

import (
	"fmt"

	"main/calculator"
)

func main() {
	a, b := 5.0, 3.0

	calc := calculator.NewCalculator(calculator.AddStrategy{})
	run(calc, a, b, "+")

	calc.SetStrategy(calculator.SubtractStrategy{})
	run(calc, a, b, "-")

	calc.SetStrategy(calculator.MultiplyStrategy{})
	run(calc, a, b, "*")

	calc.SetStrategy(calculator.DivideStrategy{})
	run(calc, a, b, "/")

	run(calc, a, 0, "/")
}

func run(calc *calculator.Calculator, a, b float64, symbol string) {
	result, err := calc.Calculate(a, b)
	if err != nil {
		fmt.Printf("%v %s %v = error: %v\n", a, symbol, b, err)
		return
	}
	fmt.Printf("%v %s %v = %v\n", a, symbol, b, result)
}
