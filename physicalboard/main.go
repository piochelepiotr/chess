package physicalboard

import (
	"fmt"

	"github.com/piochelepiotr/chess/chessboard"
)

// PhysicalBoard is the board that represents the connection to the arduino
type PhysicalBoard struct {
	chessboard *chessboard.Chessboard
}

// MovePiece moves a piece on the physical board
func (b *PhysicalBoard) MovePiece(startPosition, endPosition chessboard.Position) {
	// Compute moves
	moves := b.chessboard.FindPath(startPosition, endPosition)
	// Break Ls
	movesWithoutL := breakLs(moves, b.chessboard)
	// Groups moves
	groupedMoves := groupMoves(movesWithoutL, b.chessboard)
	fmt.Println(groupedMoves)
	// Compute motor moves (with back and forth), with electromagnet on / off
	// Sends moves to Arduino
}
