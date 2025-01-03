package main

import (
	"bufio"
	"os"
	"testing"
)

func TestRotateLeft(t *testing.T) {
	cases := []struct {
		original Direction
		turned   Direction
	}{

		{
			original: UP,
			turned:   LEFT,
		},
		{
			original: LEFT,
			turned:   DOWN,
		},
		{
			original: DOWN,
			turned:   RIGHT,
		},
		{
			original: RIGHT,
			turned:   UP,
		},
	}

	for _, test := range cases {
		turned := test.original.rotateLeft()
		if test.turned != turned {
			t.Fatalf("%v + left => %v, isntead got=%v", test.original.string(), DOWN.string(), test.turned.string())
		}
	}
}

func TestMapParse(t *testing.T) {
	filePath := "./example.txt"
	file, err := os.Open(filePath)
	if err != nil {
		t.Fatalf("Cannot read example file %s", filePath)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	m := NewMap(lines)
	if m.xWidth != 34 {
		t.Fatalf("Example map width should be 34, instead got=%v", m.xWidth)
	}
	if m.yWidth != 12 {
		t.Fatalf("Example map width should be 12, instead got=%v", m.yWidth)
	}
	if len(m.holes) != 4 {
		t.Fatalf("Example map holes amount should be 4, instead got=%v", len(m.holes))
	}
	if len(m.cockroaches) != 10 {
		t.Fatalf("Example map cockroaches amount should be 10, instead got=%v", len(m.cockroaches))
	}
}
