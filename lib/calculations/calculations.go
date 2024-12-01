package calculations

import "math"

func DiffBetweenInts(a int, b int) int {
	difference := math.Abs(float64(a) - float64(b))
	return int(difference)
}

func AddNumbersInSlice(numberslice []int) int {
	runningTotal := 0
	for _, i := range numberslice {
		runningTotal = runningTotal + i
	}
	return runningTotal
}
