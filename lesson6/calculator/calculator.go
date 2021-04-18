package calculator

import (
	"errors"
	"fmt"
)

// Operand type
type Operand struct {
	Value float64
}

// Operator type
type Operator struct {
	Value string
}

// Validate Operator value
func (operator Operator) Validate() error {
	operators := []string{"+", "-", "*", "/"}

	for _, value := range operators {
		if value == operator.Value {
			return nil
		}
	}

	return errors.New("unknown operator")
}

// Expression type
type Expression struct {
	FirstOperand, SecondOperand Operand
	Operator                    Operator
}

// Validate Expression type
func (expr Expression) Validate() []error {
	var errorList []error

	if operatorError := expr.Operator.Validate(); operatorError != nil {
		errorList = append(errorList, operatorError)
	}

	return errorList
}

// Sum implements addition between 2 numbers
func (expr Expression) Sum() (float64, error) {
	return (expr.FirstOperand.Value + expr.SecondOperand.Value), nil
}

// Subtract implements subtraction between 2 numbers
func (expr Expression) Subtract() (float64, error) {
	return (expr.FirstOperand.Value - expr.SecondOperand.Value), nil
}

// Multiply multiplication of two numbers
func (expr Expression) Multiply() (float64, error) {
	return (expr.FirstOperand.Value * expr.SecondOperand.Value), nil
}

// Divide implements division of a number by a number
func (expr Expression) Divide() (float64, error) {
	if expr.SecondOperand.Value == 0 {
		return 0, errors.New("division by zero")
	}

	return (expr.FirstOperand.Value / expr.SecondOperand.Value), nil
}

// Calculate performs arithmetic operations between two numbers.
func (expr Expression) Calculate() (float64, error) {
	result, resultError := calc(expr)
	if resultError != nil {
		return 0, resultError
	}

	fmt.Printf("%.2f %s %.2f = %.2f\n", expr.FirstOperand.Value, expr.Operator.Value, expr.SecondOperand.Value, result)
	return result, nil
}

func calc(expr Expression) (float64, error) {
	var operatorMap = map[string]func() (float64, error){
		"+": expr.Sum,
		"-": expr.Subtract,
		"*": expr.Multiply,
		"/": expr.Divide,
	}

	if expressionFunc, ok := operatorMap[expr.Operator.Value]; ok {
		return expressionFunc()
	}

	return 0, errors.New("command not implemented")
}
