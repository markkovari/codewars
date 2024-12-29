package main

import "testing"

func TestMillipe(t *testing.T) {

	first := Millipede([]string{"excavate", "endure", "screen", "desire", "theater", "excess", "night"})
	second := Millipede([]string{"trade", "pole", "view", "grave", "ladder", "mushroom", "president"})

	if !first {
		t.Fatalf("fist should be true, got=%t", first)
	}

	if second {
		t.Fatalf("second should be false, got=%t", second)
	}
}
