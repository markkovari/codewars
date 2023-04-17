package main

import "testing"

func TestFirstNonRepeating(t *testing.T) {
	t.Run("returns the first non-repeating character", func(t *testing.T) {
		got := FirstNonRepeating("a")
		want := "a"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("returns the first non-repeating character asd", func(t *testing.T) {
		got := FirstNonRepeating("stress")
		want := "t"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
