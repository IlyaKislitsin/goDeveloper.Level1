package uintcalculator

import (
	"errors"
	"fmt"
)

// UintExpression type
type UintExpression struct {
	FirstOperand, SecondOperand uint64
	Operator                    string
}

// Validate UintExpression type
func (expr UintExpression) Validate() []error {
	var errorList []error

	if operatorError := operatorValidate(expr.Operator); operatorError != nil {
		errorList = append(errorList, operatorError)
	}

	return errorList
}

// Sum implements addition between 2 numbers
func (expr UintExpression) Sum() (interface{}, error) {
	return (expr.FirstOperand + expr.SecondOperand), nil
}

// Subtract implements subtraction between 2 numbers
func (expr UintExpression) Subtract() (interface{}, error) {
	return (expr.FirstOperand - expr.SecondOperand), nil
}

// Multiply multiplication of two numbers
func (expr UintExpression) Multiply() (interface{}, error) {
	return (expr.FirstOperand * expr.SecondOperand), nil
}

// Divide implements division of a number by a number
func (expr UintExpression) Divide() (interface{}, error) {
	if expr.SecondOperand == 0 {
		return 0, errors.New("division by zero")
	}

	return (expr.FirstOperand / expr.SecondOperand), nil
}

// Calculate performs arithmetic operations between two numbers.
func (expr UintExpression) Calculate() (interface{}, error) {
	result, resultError := calc(expr)
	if resultError != nil {
		return 0, resultError
	}

	fmt.Printf("%d %s %d = %d\n", expr.FirstOperand, expr.Operator, expr.SecondOperand, result)
	return result, nil
}

func operatorValidate(operator string) error {
	operators := []string{"+", "-", "*", "/"}

	for _, value := range operators {
		if value == operator {
			return nil
		}
	}

	return errors.New("unknown operator")
}

func calc(expr UintExpression) (interface{}, error) {
	var operatorMap = map[string]func() (interface{}, error){
		"+": expr.Sum,
		"-": expr.Subtract,
		"*": expr.Multiply,
		"/": expr.Divide,
	}

	if expressionFunc, ok := operatorMap[expr.Operator]; ok {
		return expressionFunc()
	}

	return 0, errors.New("command not implemented")
}
