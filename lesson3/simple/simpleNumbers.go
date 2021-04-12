package simple

import "fmt"

// Find simple numbers to limit
func Find(limit int64) ([]int64, error) {
	if limit < 2 {
		return nil, fmt.Errorf("limit must be >= 2")
	}

	const startNumber int64 = 2

	numbers := make([]bool, limit+1)

	for i := startNumber; i <= limit; i++ {
		if numbers[i] == true {
			continue
		}

		for j := i + i; j <= limit; j += i {
			if numbers[j] == true {
				continue
			}

			numbers[j] = true
		}
	}

	simpleNumbers := []int64{}

	for index, isComposite := range numbers {
		if index < 2 {
			continue
		}
		if !isComposite {
			simpleNumbers = append(simpleNumbers, int64(index))
		}
	}

	return simpleNumbers, nil
}
