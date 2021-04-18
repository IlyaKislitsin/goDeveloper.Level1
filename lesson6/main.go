package main

import (
	"fmt"
	"strconv"

	"github.com/IlyaKislitsin/goDeveloper.Level1/lesson6/calculator"
)

func main() {
	fmt.Println("Введите выражение. Формат: num1 operator( +, -, *, /) num2")

	var num1, operator, num2 string
	fmt.Scan(&num1, &operator, &num2)

	operand1, err := strconv.ParseFloat(num1, 64)
	if err != nil {
		fmt.Printf("error: num1 is not a number")
		return
	}

	operand2, err := strconv.ParseFloat(num2, 64)
	if err != nil {
		fmt.Printf("error: num2 is not a number")
		return
	}

	expression := calculator.Expression{
		FirstOperand: calculator.Operand{
			Value: operand1,
		},
		SecondOperand: calculator.Operand{
			Value: operand2,
		},
		Operator: calculator.Operator{
			Value: operator,
		},
	}

	if validateErrors := expression.Validate(); validateErrors != nil {
		fmt.Println("errors: ", validateErrors)
		return
	}

	res, calculateError := expression.Calculate()
	if calculateError != nil {
		fmt.Println("error: ", calculateError)
		return
	}
	fmt.Printf("result: %.2f", res)
}
