package main

import "testing"

func TestHelloEmpty(t *testing.T) {
	root := DigitalRoot(121)
	if root != 4 {
		t.Fatalf(`DigitalRoot 121 should be 4, got %d`, root)
	}
}
