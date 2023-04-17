package main

func Permutations(s string) []string {
	freq := make(map[rune]int)
	for _, r := range s {
		freq[r]++
	}
	return permute(freq, "", len(s))
}

func permute(freq map[rune]int, prefix string, remaining int) []string {
	if remaining == 0 {
		return []string{prefix}
	}
	result := make([]string, 0)
	for r, count := range freq {
		if count > 0 {
			freq[r]--
			result = append(result, permute(freq, prefix+string(r), remaining-1)...)
			freq[r]++
		}
	}
	return result
}
