package chessboard

import (
	"fmt"

	"github.com/RyanCarrier/dijkstra"
)

func addStraght(graph *dijkstra.Graph, board *Chessboard, p1, p2 Position) {
	if board.FreePosition(p1) && board.FreePosition(p2) {
		graph.AddArc(p1.getIndex(), p2.getIndex(), 2)
	}
}

func addDiag(graph *dijkstra.Graph, board *Chessboard, p1, p2 Position) {
	if board.FreePosition(p1) && board.FreePosition(p2) {
		graph.AddArc(p1.getIndex(), p2.getIndex(), 3)
	}
}

func addL(graph *dijkstra.Graph, board *Chessboard, p1, p2 Position) {
	if board.FreePosition(p1) && board.FreePosition(p2) {
		graph.AddArc(p1.getIndex(), p2.getIndex(), 100)
	}
}

// FindPath finds the path from start to end, by pushing the minimum number of pieces
func (board *Chessboard) FindPath(start Position, end Position) []Move {
	graph := dijkstra.NewGraph()
	squares := getSquares()
	for _, square := range squares {
		graph.AddVertex(square.getIndex())
	}
	for _, square := range squares {
		addStraght(graph, board, square, square.Add(0, 1))
		addStraght(graph, board, square, square.Add(0, -1))
		addStraght(graph, board, square, square.Add(1, 0))
		addStraght(graph, board, square, square.Add(-1, 0))
		addDiag(graph, board, square, square.Add(1, 1))
		addDiag(graph, board, square, square.Add(-1, 1))
		addDiag(graph, board, square, square.Add(1, -1))
		addDiag(graph, board, square, square.Add(-1, -1))
		addL(graph, board, square, square.Add(1, 2))
		addL(graph, board, square, square.Add(1, -2))
		addL(graph, board, square, square.Add(-1, 2))
		addL(graph, board, square, square.Add(-1, -2))
		addL(graph, board, square, square.Add(2, 1))
		addL(graph, board, square, square.Add(2, -1))
		addL(graph, board, square, square.Add(-2, 1))
		addL(graph, board, square, square.Add(-2, -1))
	}
	path, err := graph.Shortest(start.getIndex(), end.getIndex())
	if err != nil {
		fmt.Println("No path available")
		return nil
	}
	fmt.Println(path.Distance)
	for _, i := range path.Path {
		fmt.Println(fromIndex(i).toName())
	}

	moves := make([]Move, 0)
	for i := 0; i < len(path.Path)-1; i++ {
		startPosition := fromIndex(path.Path[i])
		endPosition := fromIndex(path.Path[i+1])
		moves = append(moves, BuildMove(startPosition, endPosition))
	}
	return moves
}
