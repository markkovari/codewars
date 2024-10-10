package main

import (
	"testing"
)

func TestFindEmirp(t *testing.T) {
	tests := []struct {
		n        int
		expected [3]int
	}{
		{10, [3]int{0, 0, 0}},
		{50, [3]int{4, 37, 98}},
		{100, [3]int{8, 97, 418}},
	}
	for _, test := range tests {
		if result := FindEmirp(test.n); result != test.expected {
			t.Errorf("FindEmirp(%d) = %v, expected %v", test.n, result, test.expected)
		}
	}
}

func TestReverse(t *testing.T) {
	tests := []struct {
		n        int
		expected int
	}{
		{10, 1},
		{50, 5},
		{100, 1},
		{55, 55},
	}
	for _, test := range tests {
		if result := reverse(test.n); result != test.expected {
			t.Errorf("reverse(%d) = %d, expected %d", test.n, result, test.expected)
		}
	}
}

func TestIsPrime(t *testing.T) {

	tests := []struct {
		n        int
		expected bool
	}{
		{1, false},
		{2, true},
		{3, true},
		{4, false},
		{5, true},
		{6, false},
		{11, true},
		{12, false},
	}
	for _, test := range tests {
		if result := isPrime(test.n); result != test.expected {
			t.Errorf("isPrime(%d) = %t, expected %t", test.n, result, test.expected)
		}
	}
}
