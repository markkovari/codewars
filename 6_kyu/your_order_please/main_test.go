package main

import "testing"

func TestOrder(t *testing.T) {
	tests := []struct {
		sentence string
		expected string
	}{
		{"", ""},
		{"is2 Thi1s T4est 3a", "Thi1s is2 3a T4est"},
		{"4of Fo1r pe6ople g3ood th5e the2", "Fo1r the2 g3ood 4of th5e pe6ople"},
	}

	for _, test := range tests {
		if got := Order(test.sentence); got != test.expected {
			t.Errorf("Order(%q) = %q; want %q", test.sentence, got, test.expected)
		}
	}

}

func TestExtractNumberAndWord(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		number   int
	}{
		{"", "", 0},
		{"is2", "is", 2},
		{"Thi1s", "This", 1},
		{"T4est", "Test", 4},
		{"3a", "a", 3},
		{"4of", "of", 4},
		{"Fo1r", "For", 1},
		{"pe6ople", "people", 6},
		{"g3ood", "good", 3},
		{"th5e", "the", 5},
		{"the2", "the", 2},
	}

	for _, test := range tests {
		if word, number := ExtractNumberAndWord(test.input); word != test.expected || number != test.number {
			t.Errorf("ExtractNumberAndWord(%q) = (%q, %d); want (%q, %d)", test.input, word, number, test.expected, test.number)
		}
	}
}
