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

func TestPartitions(t *testing.T) {
	simple := []int{1, 2, 3}
	err, before, elem, after := PartitionWithElement(simple, 1)
	if err != nil {
		t.Fatalf("Error should not be present here %v", simple)
	}
	if len(before) != 1 {
		t.Fatalf("Simple before should be one element, instead it has  %d elements, %v", len(before), before)
	}
	if before[0] != 1 {
		t.Fatalf("Simple before only element, should be %d, instead got %d", 1, before[0])
	}

	if elem != 2 {
		t.Fatalf("Simple only element, should be %d, instead got %d", 2, elem)
	}

	if len(after) != 1 {
		t.Fatal("Simple after should be one element")
	}
	if after[0] != 1 {
		t.Fatalf("Simple after only element, should be %d, instead got %d", 3, after[0])
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
