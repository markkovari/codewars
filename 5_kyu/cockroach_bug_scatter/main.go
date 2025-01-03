package main

import (
	"strconv"
	"strings"
)

type Direction int
type Position struct {
	x, y int
}

const (
	UP Direction = iota
	LEFT
	DOWN
	RIGHT
)

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

func (d Direction) rotateLeft() Direction {
	return Direction((d + 1) % 4)
}

func (d Direction) string() string {
	switch d {
	case LEFT:
		return "L"
	case RIGHT:
		return "R"
	case UP:
		return "U"
	case DOWN:
		return "D"
	default:
		return "HEH?"
	}
}

type Cockroach struct {
	direction Direction
	foundHole bool
	Position
}

var moveMap map[Direction][2]int = make(map[Direction][2]int)

func init() {
	moveMap[UP] = [2]int{0, -1}
	moveMap[DOWN] = [2]int{0, 1}
	moveMap[LEFT] = [2]int{-1, 0}
	moveMap[RIGHT] = [2]int{1, 0}
}

func (c Cockroach) move(fields Map) {
	delta := moveMap[c.direction]
	deltaX := delta[0]
	deltaY := delta[1]

	if c.Position.x+deltaX >= fields.xWidth || c.Position.x+deltaX < 1 {
		c.direction.rotateLeft()
	} else if c.Position.y+deltaY >= fields.yWidth || c.Position.y+deltaY < 1 {
		c.direction.rotateLeft()
	} else if fields.anyAtPos(deltaX, deltaY) {
		// NOOP
		// NOTE: btw this is pretty unclear base on the description tho, we will see
	} else {
		c.Position.x += deltaX
		c.Position.y += deltaY
	}
}

type Hole struct {
	id int
	Position
}

type Map struct {
	xWidth, yWidth int
	cockroaches    []Cockroach
	holes          []Hole
}

func (m Map) anyAtPos(x, y int) bool {
	for _, c := range m.cockroaches {
		if c.Position.x == x && c.Position.y == y {
			return true
		}
	}
	return false
}

func NewMap(lines []string) Map {
	m := Map{}
	holes := make([]Hole, 0)
	cockroaches := make([]Cockroach, 0)
	//presume all lines are the same length
	m.xWidth = len(lines[0])
	m.yWidth = len(lines)

	for x, line := range lines {
		for y, char := range strings.Split(line, "") {
			// Wall
			if char == "|" || char == "-" || char == "+" || char == " " {
				continue
			}
			// Holes
			holeIndex, err := strconv.ParseInt(char, 8, 10)
			if err != nil {
				cockroaches = append(cockroaches, Cockroach{
					direction: NewDirection(char),
					foundHole: false,
					Position: Position{
						x,
						y,
					},
				})
				continue
			}
			holes = append(holes, Hole{id: int(holeIndex), Position: Position{x, y}})
		}
	}
	m.holes = holes
	m.cockroaches = cockroaches
	return m
}

func main() {
	println("hello cockroach_bug_scatter")
}
