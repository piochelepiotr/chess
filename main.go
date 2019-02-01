package main

import "github.com/piochelepiotr/chess/chessboard"

func main() {
	board := chessboard.CreateChessboard()
	board.FindPath(chessboard.NameToPosition("G3"), chessboard.NameToPosition("B6"))
}
