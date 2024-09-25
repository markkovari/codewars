package main

func WordsToMarks(s string) int {
	var sum int
	for _, c := range s {
		sum += int(c) - 96
	}
	return sum
}

func main() {
	println(WordsToMarks("attitude")) // 100
	// test your function here
}
