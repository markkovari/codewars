package kata

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func Disemvowel(comment string) string {
	VOWELS := []string{"a", "e", "i", "o", "u", "A", "E", "I", "O", "U"}
	result := ""
	for _, c := range comment {
		if !contains(VOWELS, string(c)) {
			result += string(c)
		}
	}
	return result
}
