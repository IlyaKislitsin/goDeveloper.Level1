package sort

// ByInserts implements insertion sort.
func ByInserts(slice ...int64) []int64 {
	sortedSlice := []int64{}

	for _, currentNumber := range slice {
		inserted := false

		for index, number := range sortedSlice {
			if currentNumber <= number {
				sortedSlice = append(sortedSlice, 0)
				copy(sortedSlice[index+1:], sortedSlice[index:])
				sortedSlice[index] = currentNumber
				inserted = true
				break
			}
		}

		if inserted {
			continue
		}

		sortedSlice = append(sortedSlice, currentNumber)
	}

	return sortedSlice
}
