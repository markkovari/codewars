package main

func DigitalRoot(n int) int {
	if n > 9 {
		return DigitalRoot(n/10 + n%10)
	}
	return n
}
