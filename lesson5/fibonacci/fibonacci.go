package fibonacci

// FindRecursive searches for a given fibonacci number using recursion.
func FindRecursive(indexNumber uint64) uint64 {
	if indexNumber == 0 || indexNumber == 1 {
		return 0
	}

	if indexNumber == 2 || indexNumber == 3 {
		return 1
	}

	return FindRecursive(indexNumber-1) + FindRecursive(indexNumber-2)
}

// Find searches for a given fibonacci number.
func Find(indexNumber uint64) uint64 {
	fibonacciNumbers := make(map[uint64]uint64)
	fibonacciNumbers[1] = 0
	fibonacciNumbers[2] = 1
	fibonacciNumbers[3] = 1

	if _, ok := fibonacciNumbers[indexNumber]; ok {
		return fibonacciNumbers[indexNumber]
	}

	for i := uint64(4); i <= indexNumber; i++ {
		fibonacciNumbers[i] = fibonacciNumbers[i-1] + fibonacciNumbers[i-2]
	}

	return fibonacciNumbers[indexNumber]
}
