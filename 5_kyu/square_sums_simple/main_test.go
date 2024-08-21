package main

import "testing"

func TestSquareIncremental(t *testing.T) {
	for i := range []int{1, 2, 3} {
		small := SquareSumsRow(i)
		if small != nil {
			t.Fatalf("Two elements cannot be solved with squares, instead got %+v", small)
		}
	}

}

func TestIsSquare(t *testing.T) {

	notSquares := []int{1, 2, 3, 5, 6, 11, 101}
	squares := []int{4, 9, 16, 25, 36, 49, 81, 100}

	for _, notSquare := range notSquares {
		if IsSqrt(notSquare) {
			t.Fatalf("%d should not be square", notSquare)
		}
	}

	for _, square := range squares {
		if !IsSqrt(square) {
			t.Fatalf("%d should  be square", square)
		}
	}
}

func testOnceInArray(ar []int, t *testing.T) {
	if ar == nil {
		return
	}
	if len(ar) == 0 {
		return
	}
	if len(ar) == 1 {
		return
	}
	lookup := make(map[int]int)
	for _, elem := range ar {
		lookup[elem] += 1
		if lookup[elem] == 2 {
			t.Fatalf("Array contains multiple times %d", elem)
		}
	}
}
