package main

func ArrayDiff(a, b []int) []int {
	var result []int

	for _, num := range a {
		// If num is not present in b, append it to the result
		if !contains(b, num) {
			result = append(result, num)
		}
	}

	return result
}

// Helper function to check if a slice contains a specific value
func contains(slice []int, val int) bool {
	for _, num := range slice {
		if num == val {
			return true
		}
	}
	return false
}
