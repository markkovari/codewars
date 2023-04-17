package kata

import (
	"testing"
)

func TestSolution(t *testing.T) {

	tests := []struct {
		name string
		list []int
		want string
	}{
		{"Example Tests", []int{-6, -3, -2, -1, 0, 1, 3, 4, 5, 7, 8, 9, 10, 11, 14, 15, 17, 18, 19, 20}, "-6,-3-1,3-5,7-11,14,15,17-20"},
		{"Example Tests", []int{-3, -2, -1, 2, 10, 15, 16, 18, 19, 20}, "-3--1,2,10,15,16,18-20"},
		{"Example Tests", []int{-3, -2, -1, 2, 10, 15, 16, 18, 19, 20}, "-3--1,2,10,15,16,18-20"},
		{"Example Tests", []int{-3, -2, -1, 2, 10, 15, 16, 18, 19, 20}, "-3--1,2,10,15,16,18-20"},
		{"Example Tests", []int{-3, -2, -1, 2, 10, 15, 16, 18, 19, 20}, "-3--1,2,10,15,16,18-20"},
		{"Example Tests", []int{-3, -2, -1, 2, 10, 15, 16, 18, 19, 20}, "-3--1,2,10,15,16,18-20"},
		{"Example Tests", []int{-3, -2, -1, 2, 10, 15, 16, 18, 19, 20}, "-3--1,2,10,15,16,18-20"},
		{"Example Tests", []int{-3, -2, -1, 2, 10, 15, 16, 18, 19, 20}, "-3--1,2,10,15,16,18-20"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.list); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
