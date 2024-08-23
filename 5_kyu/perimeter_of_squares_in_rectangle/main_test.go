package main

import "testing"

func TestPerimeter(t *testing.T) {

	tests := []struct {
		n   int
		exp int
	}{
		{5, 80},
		{7, 216},
		{20, 114624},
		{30, 14098308},
	}
	for _, test := range tests {
		if got := Perimeter(test.n); got != test.exp {
			t.Errorf("Perimeter(%d) = %d; want %d", test.n, got, test.exp)
		}
	}
}
