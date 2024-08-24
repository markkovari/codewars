package main

import "testing"

func TestDecompose(t *testing.T) {

	tests := []struct {
		input string
		want  []string
	}{
		{"21/23", []string{"1/2", "1/3", "1/13", "1/359", "1/644046"}},
		{"12/4", []string{"3"}},
		{"0.66", []string{"1/2", "1/7", "1/59", "1/5163", "1/53307975"}},
		{"0", []string{}},
		{"1", []string{"1"}},
	}
	for _, test := range tests {
		if got := Decompose(test.input); !equal(got, test.want) {
			t.Errorf("Decompose(%q) = %v, want %v", test.input, got, test.want)
		}
	}
}

func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
