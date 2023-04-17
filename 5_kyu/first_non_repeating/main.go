package main

import "unicode"

func FirstNonRepeating(str string) string {
	charCount := make(map[rune]int)
	for _, char := range str {
		charCount[unicode.ToLower(char)]++
	}
	for _, char := range str {
		if charCount[unicode.ToLower(char)] == 1 {
			return string(char)
		}
	}
	return ""
}
