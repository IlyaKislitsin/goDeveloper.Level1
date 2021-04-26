package main

import (
	"fmt"

	"github.com/IlyaKislitsin/goDeveloper.Level1/lesson7/calculator"
	floatCalculator "github.com/IlyaKislitsin/goDeveloper.Level1/lesson7/calculator/float"
	intCalculator "github.com/IlyaKislitsin/goDeveloper.Level1/lesson7/calculator/int"
	uintCalculator "github.com/IlyaKislitsin/goDeveloper.Level1/lesson7/calculator/uint"
)

func main() {
	floatCalc()
	intCalc()
	uintCalc()
}

func floatCalc() {
	fmt.Println("FloatCalc. Введите выражение. Формат: num1 operator( +, -, *, /) num2")

	floatExpr := floatCalculator.FloatExpression{}
	if _, scanErr := fmt.Scan(&floatExpr.FirstOperand, &floatExpr.Operator, &floatExpr.SecondOperand); scanErr != nil {
		fmt.Println("error: ", scanErr)
		return
	}

	if validateErrors := validate(floatExpr); validateErrors != nil {
		fmt.Println("errors: ", validateErrors)
		return
	}

	res, calculateError := calc(floatExpr)
	if calculateError != nil {
		fmt.Println("error: ", calculateError)
		return
	}
	fmt.Printf("floatCalc result: %.2f\n", res)
}

func intCalc() {
	fmt.Println("IntCalc. Введите выражение. Формат: num1 operator( +, -, *, /) num2")

	intExpr := intCalculator.IntExpression{}
	fmt.Scan(&intExpr.FirstOperand, &intExpr.Operator, &intExpr.SecondOperand)
	if _, scanErr := fmt.Scan(&intExpr.FirstOperand, &intExpr.Operator, &intExpr.SecondOperand); scanErr != nil {
		fmt.Println("error: ", scanErr)
		return
	}

	if validateErrors := validate(intExpr); validateErrors != nil {
		fmt.Println("errors: ", validateErrors)
		return
	}

	res, calculateError := calc(intExpr)
	if calculateError != nil {
		fmt.Println("error: ", calculateError)
		return
	}
	fmt.Printf("intCalc result: %d\n", res)
}

func uintCalc() {
	fmt.Println("UintCalc. Введите выражение. Формат: num1 operator( +, -, *, /) num2")

	uintExpr := uintCalculator.UintExpression{}
	fmt.Scan(&uintExpr.FirstOperand, &uintExpr.Operator, &uintExpr.SecondOperand)
	if _, scanErr := fmt.Scan(&uintExpr.FirstOperand, &uintExpr.Operator, &uintExpr.SecondOperand); scanErr != nil {
		fmt.Println("error: ", scanErr)
		return
	}

	if validateErrors := validate(uintExpr); validateErrors != nil {
		fmt.Println("errors: ", validateErrors)
		return
	}

	res, calculateError := calc(uintExpr)
	if calculateError != nil {
		fmt.Println("error: ", calculateError)
		return
	}
	fmt.Printf("uintCalc result: %d\n", res)
}

func validate(expression calculator.Calculator) []error {
	return expression.Validate()
}

func calc(expression calculator.Calculator) (interface{}, error) {
	return expression.Calculate()
}
