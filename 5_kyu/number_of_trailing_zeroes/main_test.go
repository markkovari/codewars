package main

import (
	"testing"
)

func TestZeros(t *testing.T) {
	t.Run("Test Zeros 0", func(t *testing.T) {
		got := Zeros(0)
		want := 0

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
	t.Run("Test Zeros 6", func(t *testing.T) {
		got := Zeros(6)
		want := 1

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
	t.Run("Test Zeros 30", func(t *testing.T) {
		got := Zeros(30)
		want := 7

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}
