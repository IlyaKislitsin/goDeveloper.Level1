package calculator

import (
	"fmt"
	"os"
	"strconv"
)

// Calculate performs arithmetic operations between two numbers.
func Calculate() {
	fmt.Println("Калькулятор.")
	fmt.Print("Введите первое число: ")
	a, errorA := getNumber()
	if errorA != nil {
		fmt.Println("Ошибка при получении первого числа! Невозможно произвести вычисление.")
		return
	}

	fmt.Print("Введите второе число: ")
	b, errorB := getNumber()
	if errorB != nil {
		fmt.Println("Ошибка при получении второго числа! Невозможно произвести вычисление.")
		return
	}

	fmt.Print("Введите опертор(+, -, *, /): ")
	operator, errorOperator := getOperator()
	if errorOperator != nil {
		fmt.Println("Ошибка при получении ооператора! Невозможно произвести вычисление.")
		return
	}

	result, resultError := calc(a, b, operator)
	if resultError != nil {
		fmt.Println("Ошибка вычисления!")
		return
	}
	fmt.Printf("%v %s %v = %d\n", a, operator, b, result)
}

func getNumber() (int64, error) {
	var s string
	for {
		_, scanErr := fmt.Scan(&s)
		number, err := strconv.ParseInt(s, 10, 64)
		if scanErr != nil || err != nil {
			fmt.Fprintf(os.Stderr, "Ошибка: введены некорректные данные. Попробуйте ещё раз: ")
			continue
		}
		return number, nil
	}
}

func getOperator() (string, error) {
	operators := []string{"+", "-", "*", "/"}
	var operator string
	for {
		_, scanErr := fmt.Scan(&operator)
		if scanErr != nil {
			fmt.Fprintf(os.Stderr, "Ошибка: введены некорректные данные. Попробуйте ещё раз: ")
			continue
		}

		if !findInSlice(operator, operators) {
			fmt.Print("Неизвестный опрератор. Попробуйте ещё раз: ")
			continue
		}

		return operator, nil
	}
}

func findInSlice(val string, slice []string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

func calc(num1 int64, num2 int64, operator string) (int64, error) {
	var result int64
	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if (num2 == 0) {
			return result, fmt.Errorf("на 0 делить нельзя")
		}
		result = num1 / num2
	default:
		return result, fmt.Errorf("неизвестный оператор")
	}

	return result, nil
}
