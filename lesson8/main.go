package main

import (
	"fmt"
	"strconv"

	"github.com/IlyaKislitsin/goDeveloper.Level1/lesson7/calculator"
	floatCalculator "github.com/IlyaKislitsin/goDeveloper.Level1/lesson7/calculator/float"
	"github.com/IlyaKislitsin/goDeveloper.Level1/lesson8/config"
)

func main() {
	config, err := config.New()
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	// config.Print()

	if *config.Greeting != "" {
		fmt.Println(*config.Greeting)
	}
	floatCalc(config)
}

func floatCalc(c config.Config) {
	fmt.Println("FloatCalc. Введите выражение. Формат: num1 operator( +, -, *, /) num2")

	floatExpr := floatCalculator.FloatExpression{}
	if _, scanErr := fmt.Scan(&floatExpr.FirstOperand, &floatExpr.Operator, &floatExpr.SecondOperand); scanErr != nil {
		fmt.Println("error: ", scanErr)
		return
	}

	if *c.Unsigned && (floatExpr.FirstOperand <= 0 || floatExpr.SecondOperand <= 0) {
		fmt.Println("error: operands must be unsigned")
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
	fmt.Printf("floatCalc result: %."+strconv.Itoa(int(*c.FloatPrecision))+"f\n", res)
}

func validate(expression calculator.Calculator) []error {
	return expression.Validate()
}

func calc(expression calculator.Calculator) (interface{}, error) {
	return expression.Calculate()
}
