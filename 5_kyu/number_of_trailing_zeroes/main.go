package main

func Zeros(n int) int {
	count := 0
	for i := 5; n/i >= 1; i *= 5 {
		count += n / i
	}
	return count
}
