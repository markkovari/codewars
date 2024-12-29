package main

import "testing"

func TestCypher(t *testing.T) {
	type Case struct {
		original string
		cypher   string
		result   string
	}

	tests := []Case{
		{
			original: "Ala has a cat",
			cypher:   "gaderypoluki",
			result:   "Gug hgs g cgt",
		},
	}

	for _, test := range tests {
		res := Encode(test.original, test.cypher)
		if res != test.result {
			t.Fatalf("Cannot cyher {{%v}} + {{%v}} = {{%v}}, and != {{%v}}", test.original, test.cypher, test.result, res)
		}
		rev := Decode(test.result, test.cypher)
		if rev != test.original {
			t.Fatalf("Cannot decyher {{%v}} + {{%v}} = {{%v}}, and != {{%v}}", test.original, test.cypher, test.result, rev)
		}
	}
}
