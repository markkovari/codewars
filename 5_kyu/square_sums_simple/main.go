package main

import (
	"fmt"
	"math"
)

func SquareSumsRow(n int) []int {
	if n <= 3 {
		return nil
	}
	initial := GenerateIncrementalArrayUntil(n)
	return generatePermutationsRecursive(initial, IsSqaureSumSeries)
}

func generatePermutationsRecursive(numbers []int, evaluationFunc func([]int) bool) []int {
	if len(numbers) == 0 {
		return nil
	}
	if len(numbers) == 1 {
		if evaluationFunc(numbers) {
			return numbers
		}
		return nil
	}
	varations := make(chan []int)

	for i := 0; i < len(numbers); i++ {
		before, num, after := Partition(numbers, i)
		for _, item := range Merge(before, after) {
			if PairSumIsSquare(num, item) {
				println(fmt.Sprintf("before: %+v after: %+v number: %d item: %d\n", before, after, num, item))
				varations <- []int{num, item}
			}
		}
	}
	for more := range varations {
		if evaluationFunc(more) {
			return more
		}
	}
	return nil
}

func Partition(arr []int, i int) ([]int, int, []int) {
	return arr[:i], arr[i], arr[i+1:]
}

func Merge(arr1, arr2 []int) []int {
	return append(arr1, arr2...)
}

func GenerateIncrementalArrayUntil(n int) []int {
	if n < 0 {
		return nil
	}
	if n == 0 {
		return []int{}
	}
	result := make([]int, n)
	for i := 0; i < n; i += 1 {
		result[i] = i + 1
	}
	return result
}

func IsSqaureSumSeries(arr []int) bool {
	return HasNoDuplicate(arr) && EveryPairSumIsSquare(arr)
}

func EveryPairSumIsSquare(arr []int) bool {
	for i := 0; i < len(arr)-1; i += 1 {
		if !PairSumIsSquare(arr[i], arr[i+1]) {
			return false
		}
	}

	return true
}

func PairSumIsSquare(a, b int) bool {
	return IsSqrt(a + b)
}

func IsSqrt(aNumber int) bool {
	if aNumber == 1 {
		return false
	}
	theSquare := math.Floor(math.Sqrt(float64(aNumber)))
	return math.Pow(theSquare, 2) == float64(aNumber)
}

func HasNoDuplicate(arr []int) bool {
	counter := make(map[int]int)
	for _, elem := range arr {
		_, alreadyExists := counter[elem]
		if alreadyExists {
			return false
		} else {
			counter[elem] = 1
		}
	}
	return true
}

func main() {
	println("HELLO YO")
}
