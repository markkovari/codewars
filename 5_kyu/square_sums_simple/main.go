package main

import (
	"fmt"
	"math"
)

func SquareSumsRow(n int) []int {
	if n <= 2 {
		return nil
	}
	return nil
}

func PartitionWithElement(arr []int, posintion int) (error, []int, int, []int) {
	if posintion < 0 || posintion >= len(arr) {
		return fmt.Errorf("wrong index : %d, %d elements only", posintion, len(arr)-1), nil, 0, nil
	}

	// there might be some cheeky breeky edgecases, use it with caution
	return nil, arr[:posintion], arr[posintion], arr[:posintion]
}

func IsSqrt(aNumber int) bool {
	if aNumber == 1 {
		return false
	}
	theSquare := math.Floor(math.Sqrt(float64(aNumber)))
	return math.Pow(theSquare, 2) == float64(aNumber)
}

func main() {
	println("HELLO YO")
}
