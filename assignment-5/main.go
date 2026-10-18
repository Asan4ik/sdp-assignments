package main

import (
	"fmt"
	"main/calculator"
)

func main() {
	calc := calculator.NewCalculator(calculator.AddStrategy{})
	var a float64 = 5
	var b float64 = 3
	calc.SetStrategy(calculator.AddStrategy{})
	fmt.Printf("%v + %v = %v\n", a, b, calc.Calculate(a, b))
	calc.SetStrategy(calculator.SubtractStrategy{})
	fmt.Printf("%v - %v = %v\n", a, b, calc.Calculate(a, b))
}
