package main

import "testing"

func TestOneEqOne(t *testing.T) {
	type testCase struct {
		h     int
		m     int
		s     int
		since int
	}
	tests := []testCase{

		{h: 0, m: 1, s: 1, since: 61000},
		{h: 1, m: 1, s: 1, since: 3661000},
		{h: 0, m: 0, s: 0, since: 0},
		{h: 1, m: 0, s: 1, since: 3601000},
		{h: 1, m: 0, s: 0, since: 3600000},
	}
	for _, test := range tests {
		if test.since != Past(test.h, test.m, test.s) {

			t.Fatalf("past is not %d milliseconds, got=%d with %d %d %d", test.since, Past(test.h,test.m,test.s), test.h, test.m, test.s)
		}

	}
}
