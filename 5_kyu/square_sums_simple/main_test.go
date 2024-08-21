package main

import "testing"

func TestSquareIncremental(t *testing.T) {
	for i := range []int{1, 2, 3} {
		small := SquareSumsRow(i)
		if small != nil {
			t.Fatalf("Two elements cannot be solved with squares, instead got %+v", small)
		}
	}

	given := SquareSumsRow(15)
	if given == nil {
		t.Fatalf("Given should Not be null :( %v", given)
	}
}

func TestGenerateIncrementalArrayUntil(t *testing.T) {
	res := GenerateIncrementalArrayUntil(1)
	if len(res) != 1 {
		t.Fatalf("Arr legth should be %d, got %d", 1, len(res))
	}
	if res[0] != 1 {
		t.Fatalf("Arr first element should be %d, got %d", 1, res[0])
	}

	res = GenerateIncrementalArrayUntil(3)
	if len(res) != 3 {
		t.Fatalf("Arr legth should be %d, got %d", 3, len(res))
	}
	if res[0] != 1 {
		t.Fatalf("Arr first element should be %d, got %d", 1, res[0])
	}
	if res[1] != 2 {
		t.Fatalf("Arr first element should be %d, got %d", 2, res[1])
	}
	if res[2] != 3 {
		t.Fatalf("Arr first element should be %d, got %d", 3, res[2])
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

	withError := []int{}
	err, before, elem, after = PartitionWithElement(withError, 420)
	if err == nil {
		t.Fatalf("Should be error since overindexing %v", withError)
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
