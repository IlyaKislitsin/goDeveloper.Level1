package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	area, _ := getArea()
	r := math.Sqrt(area / math.Pi)

	fmt.Printf("Диаметр окружности равна: %.2f\n", 2*r)
	fmt.Printf("Длина окружности равна: %.2f", 2*math.Pi*r)
}

func getArea() (float64, error) {
	var s string
	fmt.Print("Найдём диаметр и длину окружности по заданной площади круга.\n", "Введите площадь круга: ")
	for {
		_, scanErr := fmt.Scan(&s)
		area, err := strconv.ParseInt(s, 10, 64)
		if scanErr != nil || err != nil || area <= 0 {
			fmt.Fprintf(os.Stderr, "Ошибка: введены некорректные данные. Попробуйте ещё раз: ")
			continue
		}
		return float64(area), nil
	}
}
