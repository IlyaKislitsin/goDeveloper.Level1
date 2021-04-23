package intcalculator

import (
	"errors"
	"fmt"
)

// IntExpression type
type IntExpression struct {
	FirstOperand, SecondOperand int64
	Operator                    string
}

// Validate IntExpression type
func (expr IntExpression) Validate() []error {
	var errorList []error

	if operatorError := operatorValidate(expr.Operator); operatorError != nil {
		errorList = append(errorList, operatorError)
	}

	return errorList
}

// Sum implements addition between 2 numbers
func (expr IntExpression) Sum() (interface{}, error) {
	return (expr.FirstOperand + expr.SecondOperand), nil
}

// Subtract implements subtraction between 2 numbers
func (expr IntExpression) Subtract() (interface{}, error) {
	return (expr.FirstOperand - expr.SecondOperand), nil
}

// Multiply multiplication of two numbers
func (expr IntExpression) Multiply() (interface{}, error) {
	return (expr.FirstOperand * expr.SecondOperand), nil
}

// Divide implements division of a number by a number
func (expr IntExpression) Divide() (interface{}, error) {
	if expr.SecondOperand == 0 {
		return 0, errors.New("division by zero")
	}

	return (expr.FirstOperand / expr.SecondOperand), nil
}

// Calculate performs arithmetic operations between two numbers.
func (expr IntExpression) Calculate() (interface{}, error) {
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

func calc(expr IntExpression) (interface{}, error) {
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
