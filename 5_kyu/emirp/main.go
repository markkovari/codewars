package main

import "math"

var primes = make(map[int]bool)
var emirps = make(map[int]bool)

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	if val, ok := primes[n]; ok {
		return val
	}
	for i := 2; i < int(math.Sqrt(float64(n)))+1; i++ {
		if n%i == 0 {
			primes[n] = false
			return false
		}
	}
	primes[n] = true
	return true
}

func isPalindrome(n int) bool {
	return n == reverse(n)
}

func reverse(n int) int {
	reversed := 0
	for n > 0 {
		reversed = reversed*10 + n%10
		n /= 10
	}
	return reversed
}

func FindEmirp(n int) [3]int {
	biggestEmirp := 0
	numberOfEmirsBelowN := 0
	sumOfEmirsBelowN := 0
	for i := 1; i <= n; i++ {
		if isPrime(i) {
			if isPrime(reverse(i)) {
				if i != reverse(i) {
					numberOfEmirsBelowN++
					sumOfEmirsBelowN += i
					if i > biggestEmirp {
						biggestEmirp = i
					}
				}
			}
		}
	}
	return [3]int{
		numberOfEmirsBelowN,
		biggestEmirp,
		sumOfEmirsBelowN,
	}
}

func main() {

}
