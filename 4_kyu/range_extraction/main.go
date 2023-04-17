package kata

import (
	"fmt"
	"strconv"
	"strings"
)

func Solution(args []int) string {
	args = append(args, 0) // Append a sentinel value to avoid index out of range errors
	stringList := []string{}
	rangeNum := []int{args[0]}

	for i := 1; i < len(args); i++ {
		if args[i] != args[i-1]+1 {
			if len(rangeNum) == 1 {
				stringList = append(stringList, strconv.Itoa(rangeNum[0]))
				rangeNum = []int{}
			} else if len(rangeNum) == 2 {
				stringList = append(stringList, strconv.Itoa(rangeNum[0]), strconv.Itoa(rangeNum[1]))
				rangeNum = []int{}
			} else if len(rangeNum) > 2 {
				stringList = append(stringList, fmt.Sprintf("%d-%d", rangeNum[0], rangeNum[len(rangeNum)-1]))
				rangeNum = []int{}
			}
		}
		rangeNum = append(rangeNum, args[i])
	}
	return strings.Join(stringList, ",")
}
