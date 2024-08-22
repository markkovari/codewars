package main

import "fmt"

type TribonacciM struct {
	memo map[int]float64
}

func (t *TribonacciM) get(n int) float64 {
	if n < 0 {
		return 0
	}

	if v, ok := t.memo[n]; ok {
		return v
	}

	t.memo[n] = t.get(n-3) + t.get(n-2) + t.get(n-1)
	return t.memo[n]
}

func Tribonacci(signature [3]float64, n int) []float64 {
	var tribonacciM TribonacciM
	tribonacciM.memo = make(map[int]float64)
	tribonacciM.memo[0] = signature[0]
	tribonacciM.memo[1] = signature[1]
	tribonacciM.memo[2] = signature[2]
	if n < 3 {
		result := make([]float64, n)
		for i := 0; i < n; i++ {
			result[i] = signature[i]
		}
		return result
	}

	tribonacciM.get(n)

	result := make([]float64, n)
	for i := 0; i < n; i++ {
		result[i] = tribonacciM.memo[i]
	}

	return result
}

func main() {
	//0, 0, 1, 1, 2, 4, 7, 13, 24, 44
	println(fmt.Sprintf("%v", Tribonacci([3]float64{0, 1, 1}, 10)))
}
