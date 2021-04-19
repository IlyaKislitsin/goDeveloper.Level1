package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/IlyaKislitsin/goDeveloper.Level1/lesson5/fibonacci"
)

func main() {
	fmt.Print("Поиск числа Фибоначи по порядковому номеру.\nВведите порядковый номер в последовательности: ")

	var in string

	_, scanErr := fmt.Scan(&in)
	if scanErr != nil {
		fmt.Fprint(os.Stderr, "Ошибка: ", scanErr)
		return
	}

	index, err := strconv.ParseUint(in, 10, 64)
	if err != nil {
		fmt.Fprint(os.Stderr, "Ошибка: ", err)
		return
	}

	// start1 := time.Now()
	fmt.Println(fibonacci.FindRecursive(index))
	// duration1 := time.Since(start1)
	// fmt.Println(duration1)

	// start2 := time.Now()
	fmt.Println(fibonacci.Find(index))
	// duration2 := time.Since(start2)
	// fmt.Println(duration2)
}
