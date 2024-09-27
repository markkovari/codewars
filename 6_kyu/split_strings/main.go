package main

import "fmt"

func Solution(str string) []string {
	result := make([]string, 0)
	for i := 0; i < len(str); i += 2 {
		first := string(str[i])
		var second string
		if i+1 >= len(str) {
			second = "_"
		} else {
			second = string(str[i+1])
		}
		currentPair := fmt.Sprintf("%s%s", first, second)
		result = append(result, currentPair)
	}
	return result
}

func main() {

}
