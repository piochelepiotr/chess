package chessboard

const (
	boardHeight = 8
	//width is 8 + 4, because we save 2 columns on each side to put the dead pieces
	boardWidth = 8 + 4
)

// Chessboard represents the board with pieces on it
type Chessboard struct {
	pieces [][]Piece
}

func (b *Chessboard) setPiece(p Position, piece Piece) {
	b.pieces[p.X][p.Y] = piece
}

func (b *Chessboard) getPiece(p Position) Piece {
	return b.pieces[p.X][p.Y]
}

// CreateChessboard inits the chessboard with the default chessboard configuration
func CreateChessboard() *Chessboard {
	pieces := make([][]Piece, boardWidth)
	for x := 0; x < boardWidth; x++ {
		pieces[x] = make([]Piece, boardHeight)
		for y := 0; y < boardHeight; y++ {
			pieces[x][y] = Piece{pawn, transparent}
		}
	}
	board := Chessboard{pieces}
	for x := 2; x < 10; x++ {
		board.setPiece(Position{x, 1}, Piece{pawn, white})
		board.setPiece(Position{x, 6}, Piece{pawn, black})
	}
	board.setPiece(NameToPosition("H1"), Piece{rook, white})
	board.setPiece(NameToPosition("G1"), Piece{knight, white})
	board.setPiece(NameToPosition("F1"), Piece{bishop, white})
	board.setPiece(NameToPosition("E1"), Piece{king, white})
	board.setPiece(NameToPosition("D1"), Piece{queen, white})
	board.setPiece(NameToPosition("C1"), Piece{bishop, white})
	board.setPiece(NameToPosition("B1"), Piece{knight, white})
	board.setPiece(NameToPosition("A1"), Piece{rook, white})

	board.setPiece(NameToPosition("H8"), Piece{rook, black})
	board.setPiece(NameToPosition("G8"), Piece{knight, black})
	board.setPiece(NameToPosition("F8"), Piece{bishop, black})
	board.setPiece(NameToPosition("E8"), Piece{king, black})
	board.setPiece(NameToPosition("D8"), Piece{queen, black})
	board.setPiece(NameToPosition("C8"), Piece{bishop, black})
	board.setPiece(NameToPosition("B8"), Piece{knight, black})
	board.setPiece(NameToPosition("A8"), Piece{rook, black})
	return &board
}

func getSquares() []Position {
	squares := make([]Position, 0)
	for x := 0; x < boardWidth; x++ {
		for y := 0; y < boardHeight; y++ {
			squares = append(squares, Position{x, y})
		}
	}
	return squares
}

// FreePosition returns true if the position
// is in the board and doesn't have a piece on it
func (b *Chessboard) FreePosition(p Position) bool {
	return positionInBoard(p) && b.getPiece(p).color == transparent
}

func positionInBoard(p Position) bool {
	return p.X >= 0 && p.Y > 0 && p.X < boardWidth && p.Y < boardHeight
}
