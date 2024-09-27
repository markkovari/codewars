package main

import (
	"testing"
)

func TestSolution(t *testing.T) {
	sol1 := Solution("abc")
	sol1[0] = "ab"
	sol1[1] = "c_"

	sol2 := Solution("dawsd")
	sol2[0] = "da"
	sol2[1] = "ws"
	sol2[2] = "d_"

	sol3 := Solution("awsaws")
	sol3[0] = "aw"
	sol3[1] = "sa"
	sol3[2] = "ws"
}
