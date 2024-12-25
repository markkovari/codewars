package main

func PrimeBefAft(n int) [2]int {
	smaller := getSmaller(n)
	bigger := getBigger(n)
	return [2]int{smaller, bigger}
}

func getSmaller(n int) int {
	current := n - 1
	for {
		if isPrime(current) {
			return current
		}
		if n == 1 {
			return -1
		}
		current -= 1
	}
}

func getBigger(n int) int {

	current := n + 1
	for {
		if isPrime(current) {
			return current
		}
		if n == 1 {
			return -1
		}
		current += 1
	}
}

func isPrime(n int) bool {

	if n < 2 {
		return false
	}
	for i := 2; i < (n/2)+1; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
