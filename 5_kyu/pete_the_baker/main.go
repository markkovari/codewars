package kata

func Cakes(recipe, available map[string]int) int {

	for k, v := range recipe {
		if available[k] < v {
			return 0
		}
	}

	min := 0
	for k, v := range recipe {
		if min == 0 {
			min = available[k] / v
		} else {
			if available[k]/v < min {
				min = available[k] / v
			}
		}
	}

	return min

}
