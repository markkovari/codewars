package main

// Assuming the lengths are at least 1
func TwoWordsMatch(first, second string) bool {
	return first[len(first)-1] == second[0]
}

func AllWordsMatch(words ...string) bool {
	for i := 0; i < len(words)-1; i++ {
		if !TwoWordsMatch(words[i], words[i+1]) {
			return false
		}
	}
	return true
}

func generatePermutations(arr []string) [][]string {
	var result [][]string
	permute(arr, 0, &result)
	return result
}

func permute(arr []string, start int, result *[][]string) {
	if start >= len(arr) {
		perm := make([]string, len(arr))
		copy(perm, arr)
		*result = append(*result, perm)
		return
	}

	for i := start; i < len(arr); i++ {
		arr[start], arr[i] = arr[i], arr[start]
		permute(arr, start+1, result)
		arr[start], arr[i] = arr[i], arr[start]
	}
}

func Millipede(words []string) bool {
	for _, perm := range generatePermutations(words) {
		if AllWordsMatch(perm...) {
			return true
		}
	}
	return false
}
