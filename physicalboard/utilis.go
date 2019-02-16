package physicalboard

import (
	"github.com/piochelepiotr/chess/chessboard"
	"github.com/piochelepiotr/chess/utils"
)

func breakLs(moves []chessboard.Move, board *chessboard.Chessboard) []chessboard.Move {
	movesWithoutL := make([]chessboard.Move, 0)
	for _, move := range moves {
		if move.IsL() {
			shift := move.Shift()
			dir1 := chessboard.CreatePosition(utils.Sign(shift.X), 0)
			dir2 := chessboard.CreatePosition(0, utils.Sign(shift.Y))
			if utils.Absolute(shift.X) == 2 {
				dir1 = chessboard.CreatePosition(0, utils.Sign(shift.Y))
				dir2 = chessboard.CreatePosition(utils.Sign(shift.X), 0)
			}
			// pos in diagonal from the start position in direction of the end postion
			diagonalPosition := move.StartPosition.Add(dir1.X, dir1.Y).Add(dir2.X, dir2.Y)
			// pos straight on the long branch of the L from the start position in direction of the end postion
			straightPosition := move.StartPosition.Add(dir2.X, dir2.Y)
			if board.FreePosition(diagonalPosition) {
				movesWithoutL = append(movesWithoutL, chessboard.BuildMove(move.StartPosition, diagonalPosition))
				movesWithoutL = append(movesWithoutL, chessboard.BuildMove(diagonalPosition, move.EndPosition))
			} else if board.FreePosition(straightPosition) {
				movesWithoutL = append(movesWithoutL, chessboard.BuildMove(move.StartPosition, straightPosition))
				movesWithoutL = append(movesWithoutL, chessboard.BuildMove(straightPosition, move.EndPosition))
			} else {
				straight2Position := move.StartPosition.Add(2*dir2.X, 2*dir2.Y)
				movesWithoutL = append(movesWithoutL, chessboard.BuildSemiMove(diagonalPosition, straight2Position, true, false))
				movesWithoutL = append(movesWithoutL, chessboard.BuildMove(move.StartPosition, diagonalPosition))
				movesWithoutL = append(movesWithoutL, chessboard.BuildMove(diagonalPosition, move.EndPosition))
				movesWithoutL = append(movesWithoutL, chessboard.BuildSemiMove(straight2Position, diagonalPosition, false, true))
			}
		} else {
			movesWithoutL = append(movesWithoutL, move)
		}
	}
	return movesWithoutL
}

func groupMoves(moves []chessboard.Move, board *chessboard.Chessboard) []chessboard.Move {
	groupedMoves := make([]chessboard.Move, 0)
	n := len(moves)
	if n == 0 {
		return groupedMoves
	}
	currentShift := moves[0].Shift()
	i := 1
	for i < n {
		if moves[i].Shift() == currentShift && groupedMoves[len(groupedMoves)-1].EndPosition == moves[i].StartPosition {
			groupedMoves[len(groupedMoves)-1].EndPosition = moves[i].EndPosition
		} else {
			currentShift = moves[i].Shift()
			groupedMoves = append(groupedMoves, moves[i])
		}
		i++
	}
	return groupedMoves
}
