package main

import (
	"testing"
)

func arrEq(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestHelloName(t *testing.T) {
	diff := ArrayDiff([]int{1, 2}, []int{1})
	if !arrEq(diff, []int{2}) {
		t.Error("Expected 1, got ", len(diff))
	}
}

func SomeOtherTests(t *testing.T) {
	diff := ArrayDiff([]int{1, 2, 2}, []int{1})
	if !arrEq(diff, []int{2, 2}) {
		t.Error("Expected 1, got ", len(diff))
	}
}

func SomeOtherTests2(t *testing.T) {
	diff := ArrayDiff([]int{1, 2, 2}, []int{2})
	if !arrEq(diff, []int{1}) {
		t.Error("Expected 1, got ", len(diff))
	}
}

func SomeOtherTests3(t *testing.T) {
	diff := ArrayDiff([]int{1, 2, 2}, []int{})
	if !arrEq(diff, []int{1, 2, 2}) {
		t.Error("Expected 1, got ", len(diff))
	}
}

func SomeOtherTests4(t *testing.T) {
	diff := ArrayDiff([]int{}, []int{1, 2})
	if !arrEq(diff, []int{}) {
		t.Error("Expected 1, got ", len(diff))
	}
}

func SomeOtherTests5(t *testing.T) {
	diff := ArrayDiff([]int{1, 2, 3}, []int{1, 2})
	if !arrEq(diff, []int{3}) {
		t.Error("Expected 1, got ", len(diff))
	}
}
