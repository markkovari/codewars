package kata

import "strconv"

const sheep = " sheep..."

func countSheep(num int) string {
	var result string
	for i := 1; i <= num; i++ {
		result += strconv.Itoa(i) + sheep
	}
	return result
}
