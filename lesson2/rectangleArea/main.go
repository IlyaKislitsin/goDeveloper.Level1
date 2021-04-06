package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Print("Найдём площадь прямоугольника.\n", "Введите длину первой стороны: ")
	a, errorA := getNumber()
	if errorA != nil {
		fmt.Println("Ошибка при получении длины первой стороны! Невозможно вычислить площадь.")
		return
	}

	fmt.Print("Введите длину второй стороны: ")
	b, errorB := getNumber()
	if errorB != nil {
		fmt.Println("Ошибка при получении длины второй стороны! Невозможно вычислить площадь.")
		return
	}

	fmt.Printf("Площадь прямоугольника равна: %d", a*b)
}

func getNumber() (int64, error) {
	var s string
	for {
		_, scanErr := fmt.Scan(&s)
		number, err := strconv.ParseInt(s, 10, 64)
		if scanErr != nil || err != nil || number <= 0 {
			fmt.Fprintf(os.Stderr, "Ошибка: введены некорректные данные. Попробуйте ещё раз: : ")
			continue
		}
		return number, nil
	}
}
