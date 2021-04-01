package main

import (
	"fmt"
	"os"
	"strconv"
)

const (
	minNumber int64 = 1
	maxNumber int64 = 999
)

func main() {
	fmt.Print("Введите число из диапазона от 1 до 999: ")

	number, _ := getNumber()
	fmt.Printf("Сотен: %v\nДесятков: %v\nЕдениц: %v", number/100, (number%100)/10, number%10)
}

func getNumber() (int64, error) {
	var s string
	for {
		_, scanErr := fmt.Scan(&s)
		number, err := strconv.ParseInt(s, 10, 64)
		if scanErr != nil || err != nil {
			fmt.Fprint(os.Stderr, "Ошибка: введены некорректные данные. Попробуйте ещё раз: ")
			continue
		}
		if number < minNumber || number > maxNumber {
			fmt.Print("Ошибка: введено число не из указанного диапазона. Попробуйте ещё раз: ")
			continue
		}
		return number, nil
	}
}
