package main

import (
	"fmt"

	"github.com/IlyaKislitsin/goDeveloper.Level1/lesson4/sort"
)

func main() {
	sorted := sort.ByInserts(45, -7654, 234, 9, -6543, 1, -591236, 93465, 123, 876)
	fmt.Println(sorted)
}
