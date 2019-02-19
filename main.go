package main

import (
	"github.com/piochelepiotr/chess/chessboard"
	"github.com/piochelepiotr/chess/physicalboard"
)

func main() {
	board := chessboard.CreateChessboard()
	physBoard := physicalboard.PhysicalBoard{
		Chessboard: board,
	}
	physBoard.MovePiece(chessboard.NameToPosition("G3"), chessboard.NameToPosition("B6"))
}
