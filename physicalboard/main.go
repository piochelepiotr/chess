package physicalboard

import (
	"fmt"

	"github.com/piochelepiotr/chess/chessboard"
)

// PhysicalBoard is the board that represents the connection to the arduino
type PhysicalBoard struct {
	Chessboard *chessboard.Chessboard
}

// MovePiece moves a piece on the physical board
func (b *PhysicalBoard) MovePiece(startPosition, endPosition chessboard.Position) {
	// Compute moves
	moves := b.Chessboard.FindPath(startPosition, endPosition)
	// Break Ls
	movesWithoutL := breakLs(moves, b.Chessboard)
	// Groups moves
	groupedMoves := groupMoves(movesWithoutL, b.Chessboard)
	// Compute motor moves (with back and forth), with electromagnet on / off
	commands := generateArduinoCommands(groupedMoves, b.Chessboard)
	// Sends moves to Arduino
	for _, command := range commands {
		fmt.Println("Sending command " + command)
		SendCommand(command)
	}
}
