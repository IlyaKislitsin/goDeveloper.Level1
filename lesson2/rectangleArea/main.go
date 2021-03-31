package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Print("Найдём площадь прямоугольника.\n", "Введите длину первой стороны: ")
	a, _ := scanNumber()

	fmt.Print("Введите длину второй стороны: ")
	b, _ := scanNumber()

	fmt.Printf("Площадь прямоугольника равна: %d", a*b)
}

func scanNumber() (uint32, error) {
	var number uint32
	for {
		_, err := fmt.Scan(&number)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Возникла ошибка: %s\nПопробуйте ещё раз: ", err)
			continue
		}
		return number, nil
	}
}
