package chessboard

import "github.com/piochelepiotr/chess/utils"

// Direction is the direction of a chess move
type Direction string

const (
	//Up move points to line 1
	Up Direction = "UP"
	//Down move
	Down Direction = "DOWN"
	//Left move
	Left Direction = "LEFT"
	//Right move
	Right Direction = "RIGHT"
	//UpRight move
	UpRight Direction = "UP_RIGHT"
	//UpLeft move
	UpLeft Direction = "UP_LEFT"
	//DownRight move
	DownRight Direction = "DOWN_RIGHT"
	//DownLeft move
	DownLeft Direction = "DOWN_LEFT"
	//VUpRightL move
	VUpRightL Direction = "V_UP_RIGHT_L"
	//VUpLeftL move
	VUpLeftL Direction = "V_UP_LEFT_L"
	//VDownRightL move
	VDownRightL Direction = "V_DOWN_RIGHT_L"
	//VDownLeftL move
	VDownLeftL Direction = "V_DOWN_LEFT_L"
	//HUpRightL move
	HUpRightL Direction = "H_UP_RIGHT_L"
	//HUpLeftL move
	HUpLeftL Direction = "H_UP_LEFT_L"
	//HDownRightL move
	HDownRightL Direction = "H_DOWN_RIGHT_L"
	//HDownLeftL move
	HDownLeftL Direction = "H_DOWN_LEFT_L"
)

// Move is a direction plus position of a move on the chessboard
type Move struct {
	Direction      Direction
	StartPosition  Position
	EndPosition    Position
	SemiFirstHalf  bool
	SemiSecondHalf bool
}

// Shift returns the shift in position that a movement implies
func (m Move) Shift() Position {
	return m.EndPosition.Add(-m.StartPosition.X, -m.StartPosition.Y)
}

// IsL returns true if a movement is an L movement
func (m Move) IsL() bool {
	s := m.Shift()
	return utils.Absolute(s.X)+utils.Absolute(s.Y) == 3
}

// BuildSemiMove builds a new move that moves the piece only half way through
func BuildSemiMove(startPosition Position, endPosition Position, firstHalf, secondHalf bool) Move {
	move := BuildMove(startPosition, endPosition)
	move.SemiFirstHalf = firstHalf
	move.SemiFirstHalf = secondHalf
	return move
}

// BuildMove builds a new move
func BuildMove(startPosition Position, endPosition Position) Move {
	//Up : y-
	//Down : y+
	//Left : x-
	//Right : x+
	direction := Up
	if endPosition == startPosition.Add(1, 1) {
		direction = DownRight
	} else if endPosition == startPosition.Add(1, -1) {
		direction = UpRight
	} else if endPosition == startPosition.Add(-1, 1) {
		direction = DownLeft
	} else if endPosition == startPosition.Add(-1, -1) {
		direction = UpLeft
	} else if endPosition == startPosition.Add(1, 0) {
		direction = Right
	} else if endPosition == startPosition.Add(0, 1) {
		direction = Down
	} else if endPosition == startPosition.Add(-1, 0) {
		direction = Left
	} else if endPosition == startPosition.Add(0, -1) {
		direction = Up
	} else if endPosition == startPosition.Add(1, 2) {
		direction = VDownRightL
	} else if endPosition == startPosition.Add(1, -2) {
		direction = VUpRightL
	} else if endPosition == startPosition.Add(-1, 2) {
		direction = VDownLeftL
	} else if endPosition == startPosition.Add(-1, -2) {
		direction = VUpLeftL
	} else if endPosition == startPosition.Add(2, 1) {
		direction = HDownRightL
	} else if endPosition == startPosition.Add(2, -1) {
		direction = HUpRightL
	} else if endPosition == startPosition.Add(-2, 1) {
		direction = HDownLeftL
	} else if endPosition == startPosition.Add(-2, -1) {
		direction = HUpLeftL
	}
	return Move{
		StartPosition: startPosition,
		EndPosition:   endPosition,
		Direction:     direction,
	}
}
