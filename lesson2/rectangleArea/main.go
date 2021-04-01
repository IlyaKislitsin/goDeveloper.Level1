package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Print("Найдём площадь прямоугольника.\n", "Введите длину первой стороны: ")
	a, _ := getNumber()

	fmt.Print("Введите длину второй стороны: ")
	b, _ := getNumber()

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
