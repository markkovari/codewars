package main

import (
	"bufio"
	"os"
	"testing"
)

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
	if m.columnMax != 34 {
		t.Fatalf("Example map width should be 34, instead got=%v", m.columnMax)
	}
	if m.rowMax != 12 {
		t.Fatalf("Example map width should be 12, instead got=%v", m.rowMax)
	}
	if len(m.holes) != 4 {
		t.Fatalf("Example map holes amount should be 4, instead got=%v", len(m.holes))
	}
	if len(m.cockroaches) != 10 {
		t.Fatalf("Example map cockroaches amount should be 10, instead got=%v", len(m.cockroaches))
	}
}

func TestProjected(t *testing.T) {
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
	m.projectMoves()
	for i, c := range m.cockroaches {
		if !c.atWalls(m) {
			t.Fatalf("Cockroach %d should be at the wall(s), instead got=%v", i, c.Position)
		}
	}
}

func TestCalculate(t *testing.T) {
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
	m.projectMoves()
	m.calculate()

}
