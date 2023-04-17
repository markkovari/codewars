package main

import (
	"testing"
)

func TestFindUniq(t *testing.T) {

	res := FindUniq([]float32{1.0, 1.0, 1.0, 2.0, 1.0, 1.0})
	if res != 2 {
		t.Errorf("Only 2 should be unique")
	}
}
func TestFindUniq2(t *testing.T) {

	res := FindUniq([]float32{0, 0, 0.55, 0, 0})
	if res != 0.55 {
		t.Errorf("Only 0.55 should be unique")
	}
}
