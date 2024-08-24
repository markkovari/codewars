package main

import (
	"strings"
	"unicode"
)

func Order(sentence string) string {
	words := strings.Split(sentence, " ")
	var wordsWithNumbers = make(map[int]string)
	if len(words) == 1 {
		_, _ = ExtractNumberAndWord(words[0])
		return words[0]
	}
	for _, word := range words {
		_, n := ExtractNumberAndWord(word)
		wordsWithNumbers[n] = word
	}

	newSentence := ""
	for i := 1; i <= len(wordsWithNumbers); i++ {
		newSentence += wordsWithNumbers[i] + " "
	}

	return newSentence[:len(newSentence)-1]
}

func ExtractNumberAndWord(input string) (string, int) {
	var word string
	var number int

	for _, r := range input {
		if unicode.IsDigit(r) {
			number = number*10 + int(r-'0')
		} else {
			word += string(r)
		}
	}

	return word, number
}
