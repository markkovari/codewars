package main

func main() {

}

func isLetter(c rune) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
}

func alphanumeric(str string) bool {
	if len(str) == 0 {
		return false
	}
	for _, c := range str {
		if !isLetter(c) {
			return false
		}
	}
	return true
}
