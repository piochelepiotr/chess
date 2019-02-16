package chessboard

import (
	"fmt"
	"strconv"
)

//Position is a position on the chessboard
type Position struct {
	X int
	Y int
}

// CreatePosition creates a new position
func CreatePosition(x, y int) Position {
	return Position{x, y}
}

func (p Position) getIndex() int {
	return boardHeight*p.X + p.Y
}

// Add adds Position{x,y} to position
func (p Position) Add(x, y int) Position {
	return Position{p.X + x, p.Y + y}
}

func (p Position) toName() string {
	const letters = "XXHGFEDCBAXX"
	return fmt.Sprintf("%c%d", letters[p.X], p.Y+1)
}

func fromIndex(index int) Position {
	return Position{index / boardHeight, index % boardHeight}
}

//NameToPosition returns the Position from the common used name
func NameToPosition(name string) Position {
	x := 0
	switch name[0] {
	case 'H':
		x = 2
	case 'G':
		x = 3
	case 'F':
		x = 4
	case 'E':
		x = 5
	case 'D':
		x = 6
	case 'C':
		x = 7
	case 'B':
		x = 8
	case 'A':
		x = 9
	}
	y, _ := strconv.Atoi(name[1:2])
	y--
	return Position{x, y}
}
