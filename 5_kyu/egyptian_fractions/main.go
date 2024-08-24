package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func solve(s string) []string {
	var res []string
	if strings.Contains(s, "/") { // the "1/2" case
		nums := strings.Split(s, "/")
		numerator, _ := strconv.ParseFloat(nums[0], 64)
		denominator, _ := strconv.ParseFloat(nums[1], 64)
		if numerator/denominator == math.Floor(numerator/denominator) {
			return []string{fmt.Sprintf("%d", int(numerator/denominator))}
		}

		return solve(fmt.Sprint(numerator / denominator))
	} else {
		num, _ := strconv.ParseFloat(s, 64)
		var i int64 = 2
		for num <= math.MaxInt64 {
			nextEval := 1.0 / float64(i)
			if nextEval < num {
				res = append(res, fmt.Sprintf("1/%d", i))
				num -= nextEval
			} else if nextEval == num {
				res = append(res, fmt.Sprintf("1/%d", i))
				break
			}
			if i > 100_000_000_000 {
				break
			}
			i++
		}
	}
	return res
}

func Decompose(s string) []string {
	if s == "1" {
		return []string{"1"}
	}
	if s == "0" {
		return []string{}
	}
	return solve(s)
}
