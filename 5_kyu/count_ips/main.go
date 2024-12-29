package main

import (
	"strconv"
	"strings"
)

func PowInts(x, n int64) int64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	y := PowInts(x, n/2)
	if n%2 == 0 {
		return y * y
	}
	return x * y * y
}

func AsIntVal(ip string) (int64, error) {
	var sum int64 = 0
	values := strings.Split(ip, ".")
	for i := len(values) - 1; i >= 0; i-- {
		num, err := strconv.ParseInt(string(values[i]), 10, 64)
		if err != nil {
			return 0, err
		}
		sum += num * PowInts(int64(256), int64(len(values)-i-1))
	}
	return sum, nil
}

func IpsBetween(start, end string) int {
	endInt, _ := AsIntVal(end)
	startInt, _ := AsIntVal(start)
	return int(endInt) - int(startInt)
}
