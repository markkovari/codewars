package main

import "testing"

func TestSquareIncremental(t *testing.T) {
	one := SquareSumsRow(1)
	testOnceInArray(one, t)
	if one[0] != 1 {
		t.Fatalf("standalon element should be 1, got=%d", one[1])
	}
	two := SquareSumsRow(2)
	if two != nil {
		t.Fatalf("Two elements cannot be solved with squares, instead got %+v", two)
	}
}

func testOnceInArray(ar []int, t *testing.T) {
	lookup := make(map[int]int)
	for _, elem := range ar {
		lookup[elem] += 1
		if lookup[elem] == 2 {
			t.Fatalf("Array contains multiple times %d", elem)
		}
	}
}
