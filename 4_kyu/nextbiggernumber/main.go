package main

import (
	"reflect"
	"sort"
	"strconv"
	"strings"
)

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

func ArePermutation(one, another string) bool {
	first := sort.StringSlice(strings.Split(one, ""))
	second := sort.StringSlice(strings.Split(another, ""))
	return reflect.DeepEqual(first, second)
}

func NextBigger(n int) int {
	if n < 10 {
		return -1
	}
	asString := strconv.FormatInt(int64(n), 10)
	for howManyLastDigits := len(asString) - 2; howManyLastDigits >= 0; howManyLastDigits-- {
		static := asString[:howManyLastDigits]
		toPermutate := asString[howManyLastDigits:]
		options := generatePermutations(strings.Split(toPermutate, ""))
		viableOptions := make([]int64, 0)
		for _, option := range options {
			asMergedString := static + strings.Join(option, "")
			asNumber, _ := strconv.ParseInt(asMergedString, 10, 64)
			if asNumber > int64(n) {
				viableOptions = append(viableOptions, asNumber)
			}
		}
		sort.Slice(viableOptions, func(i, j int) bool {
			return viableOptions[j] > viableOptions[i]
		})
		if len(viableOptions) >= 1 {
			return int(viableOptions[0])
		}
	}
	return -1
}
