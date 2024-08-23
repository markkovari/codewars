package main

func fib(n int, cache map[int]int) int {
	if n <= 1 {
		return n
	}
	if _, ok := cache[n]; !ok {
		cache[n] = fib(n-1, cache) + fib(n-2, cache)
	}
	return cache[n]
}

func sumfibUntil(n int, cache map[int]int) int {
	sum := 0
	for i := 0; i <= n; i++ {
		sum += fib(i, cache)
	}
	return sum
}

func Perimeter(n int) int {
	cache := make(map[int]int)
	cache[0] = 1
	cache[1] = 1
	return 4 * sumfibUntil(n+1, cache)
}
