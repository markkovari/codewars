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

	permutations := generatePermutationsRecursive(initial)
	for _, permutation := range permutations {
		if EveryPairSumIsSquare(permutation) {
			return permutation
		}
	}
	return nil
}

// this will be slow, we need to early return probably
func generatePermutationsRecursive(numbers []int) [][]int {
	if len(numbers) == 0 {
		return [][]int{}
	}
	if len(numbers) == 1 {
		return [][]int{{numbers[0]}}
	}

	var permutations [][]int
	for i, num := range numbers {
		remaining := append([]int{}, numbers[:i]...)
		remaining = append(remaining, numbers[i+1:]...)
		for _, p := range generatePermutationsRecursive(remaining) {
			permutations = append(permutations, append([]int{num}, p...))
		}
	}
	return permutations
}

func SearchForAvailablePermutation(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		err, left, elem, right := PartitionWithElement(arr, i)
		if err != nil {
			return nil
		}
		leftPart := append([]int{elem}, left...)
		fmt.Printf("%d, left: %v leftPart: %v right: %v\n", i, left, leftPart, right)
		return append(leftPart, SearchForAvailablePermutation(right)...)
	}
	return nil
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

func PartitionWithElement(arr []int, posintion int) (error, []int, int, []int) {
	if posintion < 0 || posintion >= len(arr) {
		return fmt.Errorf("wrong index : %d, %d elements only", posintion, len(arr)-1), nil, 0, nil
	}
	return nil, arr[:posintion], arr[posintion], arr[posintion+1:]
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
