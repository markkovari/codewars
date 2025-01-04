package main

import (
	"errors"
	"strconv"
	"strings"
)

type Direction int

const (
	UP Direction = iota
	LEFT
	DOWN
	RIGHT
)

type Position struct {
	row, column int
}

type Map struct {
	rowMax, columnMax int
	cockroaches       []*Cockroach
	holes             []Hole
}

func NewDirection(from string) Direction {
	switch from {
	case "L":
		return LEFT
	case "R":
		return RIGHT
	case "D":
		return DOWN
	default:
		return UP
	}
}

type Cockroach struct {
	direction Direction
	stopped   bool
	Position
}

func (c Cockroach) atWalls(m Map) bool {
	if c.column == 0 || c.column == m.columnMax {
		return true
	}
	if c.row == 0 || c.row == m.rowMax {
		return true
	}
	return false
}

func (c *Cockroach) project(fields Map) Position {
	switch c.direction {
	case UP:
		return Position{
			row:    c.row,
			column: 0,
		}
	case DOWN:
		return Position{
			row:    c.row,
			column: fields.columnMax,
		}
	case LEFT:
		return Position{
			row:    0,
			column: c.column,
		}
	default:
		return Position{
			row:    fields.rowMax,
			column: c.column,
		}
	}
}
func returnIfSamePosition(holes []Hole, c Cockroach) (Hole, error) {
	for _, h := range holes {
		if h.Position == c.Position {
			return h, nil
		}
	}
	return Hole{}, errors.New("no hole found")
}

type Hole struct {
	id int
	Position
}

func NewMap(lines []string) Map {
	m := Map{}
	holes := make([]Hole, 0)
	cockroaches := make([]*Cockroach, 0)
	m.columnMax = len(lines[0])
	m.rowMax = len(lines)

	for row, line := range lines {
		for column, char := range strings.Split(line, "") {
			if char == "|" || char == "-" || char == "+" || char == " " {
				continue
			}
			holeIndex, err := strconv.ParseInt(char, 8, 10)
			if err != nil {
				cockroaches = append(cockroaches, &Cockroach{
					direction: NewDirection(char),
					stopped:   false,
					Position: Position{
						row,
						column,
					},
				})
				continue
			}
			holes = append(holes, Hole{id: int(holeIndex), Position: Position{row, column}})
		}
	}
	m.holes = holes
	m.cockroaches = cockroaches
	return m
}

func (m *Map) projectMoves() {
	for _, c := range m.cockroaches {
		c.Position = c.project(*m)
		c.stopped = true
	}
}

func (m Map) calculate() [10]int {
	// counter := make(map[int]int)
	// firstHole, err := m.holeBeforeOrigo()
	// currentHoleIndex := -1 // flag to see if
	// currentCount := 0
	// if err != nil {
	// 	return [10]int{}
	// }
	// go from firsthole clockwise and count the cockroaches
	// for row := firstHole.row; row > 0; row-- {

	// }
	return [10]int{}
}

// assume there is at least one hole
func (m Map) holeBeforeOrigo() (Hole, error) {
	// Upper row
	for row := 0; row < m.rowMax; row++ {
		for _, h := range m.holes {
			if h.row == row && h.column == 0 {
				return h, nil
			}
		}
	}
	// Right column
	for column := 0; column < m.columnMax; column++ {
		for _, h := range m.holes {
			if h.column == column && h.row == m.rowMax {
				return h, nil
			}
		}
	}
	//bottom row
	for row := m.rowMax; row > 0; row-- {
		for _, h := range m.holes {
			if h.row == row && h.column == m.columnMax {
				return h, nil
			}
		}
	}
	// left column
	for column := m.columnMax; column > 0; column-- {
		for _, h := range m.holes {
			if h.column == column && h.row == 0 {
				return h, nil
			}
		}
	}
	return Hole{}, errors.New("No hole found, heh?")
}

func main() {
}
