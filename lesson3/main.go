package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/IlyaKislitsin/goDeveloper.Level1/lesson3/calculator"
)

func main() {
	fmt.Println("Урок 3. Домашняя работа. Нажмите:\n 1 - Калькулятор\n 2 - Поиск простых чисел\n 0 - Выход")
	var s string
	for {
		_, scanErr := fmt.Scan(&s)
		number, err := strconv.ParseUint(s, 10, 64)
		if scanErr != nil || err != nil {
			fmt.Fprintf(os.Stderr, "Ошибка: введены некорректные данные. Попробуйте ещё раз: : ")
			continue
		}

		switch number {
		case 1:
			calculator.Calculate()
		case 2:
		case 0:
		default:
			fmt.Print("Неизвестная команда. Попробуйте ещё раз: : ")
			continue
		}

		fmt.Println("Программа завершена.")
		break
	}
}
