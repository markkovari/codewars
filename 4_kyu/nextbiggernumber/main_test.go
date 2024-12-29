package main

import "testing"

func TestNextBigger(t *testing.T) {
	type TestCase struct {
		from int
		next int
	}
	tests := []TestCase{
		{
			from: 12,
			next: 21,
		},
		{
			from: 513,
			next: 531,
		},
		{
			from: 2017,
			next: 2071,
		},
		{
			from: 59884848459853,
			next: 59884848483559,
		},
	}

	for _, test := range tests {
		res := NextBigger(test.from)
		if res != test.next {
			t.Fatalf("Next bigger number is %d, instead got=%d", test.next, res)
		}
	}
}
