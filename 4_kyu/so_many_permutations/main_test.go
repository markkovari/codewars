package main

import (
	"reflect"
	"testing"
)

func TestPermutations(t *testing.T) {
	testCases := []struct {
		name   string
		input  string
		output []string
	}{
		{
			name:   "Empty string",
			input:  "",
			output: []string{""},
		},
		{
			name:   "Single character",
			input:  "a",
			output: []string{"a"},
		},
		{
			name:   "Two characters",
			input:  "ab",
			output: []string{"ab", "ba"},
		},
		{
			name:   "Three characters",
			input:  "abc",
			output: []string{"abc", "acb", "bac", "bca", "cba", "cab"},
		},
		{
			name:   "Duplicate characters",
			input:  "aab",
			output: []string{"aab", "aba", "baa"},
		},
		{
			name:   "All duplicate characters",
			input:  "aaa",
			output: []string{"aaa"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := Permutations(tc.input)
			if !reflect.DeepEqual(got, tc.output) {
				t.Errorf("Expected %v, but got %v", tc.output, got)
			}
		})
	}
}
