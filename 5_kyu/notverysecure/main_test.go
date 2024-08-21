package main

import "testing"

type testCase struct {
	expected bool
	input string
}

func TestAlphanumeric(t *testing.T) {
	tests := []testCase{
		{expected: false, input: ".*?"},
		{expected: true, input: "a"},
		{expected: true, input: "Mazinkaiser"},
		{expected: false, input: "hello world_"},
		{expected: true, input: "PassW0rd"},
		{expected: false, input: "     "},
		{expected: false, input: ""},
		{expected: false, input: "\n\t\n"},
		{expected: false, input: "ciao\n$$_"},
		{expected: false, input: "__ * __"},
		{expected: false, input: "&)))((("},
		{expected: true, input: "43534h56jmTHHF3k"},
	}
	for _, test := range tests {
		got := alphanumeric(test.input)
		expected := test.expected

		if got != expected {
			t.Fatalf("did not pass {%s} test got=%t want=%t", test.input, got, expected)
		}

	}
}
 

func TestIsLetter(t *testing.T) {
	truthies := "1212121212a1j2h1j2h1j2h1jh2jhakjdadoj22AAAAAbBBBB"

	for _, truthy := range truthies {
		if !isLetter(truthy) {
			t.Fatalf("should be letter {%s}", string(truthy))
		}
	}
}
