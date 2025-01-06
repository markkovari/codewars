package kata

func FindOdd(seq []int) int {
	freq := make(map[int]int)
	for _, s := range seq {
		if val, found := freq[s]; found {
			freq[s] = val + 1
		} else {
			freq[s] = 1
		}
	}
	for key, value := range freq {
		if value%2 == 1 {
			return key
		}
	}
	return 0
}
