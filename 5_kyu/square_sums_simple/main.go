package main

import "math"

func SquareSumsRow(n int) []int {
	if n <= 2 {
		return nil
	}
	return nil
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
