package physicalboard

import "github.com/piochelepiotr/chess/chessboard"

type realPosition struct {
	x float64
	y float64
}

func createRealPosition(p chessboard.Position) realPosition {
	return realPosition{
		x: float64(p.X),
		y: float64(p.Y),
	}
}

func positionBetween(p1 chessboard.Position, p2 chessboard.Position) realPosition {
	rp1 := createRealPosition(p1)
	rp2 := createRealPosition(p2)
	return realPosition{
		x: (rp1.x + rp2.x) / 2,
		y: (rp1.y + rp2.y) / 2,
	}
}

func getMoveStartPosition(m chessboard.Move) realPosition {
	if m.SemiSecondHalf {
		return positionBetween(m.StartPosition, m.EndPosition)
	}
	return createRealPosition(m.StartPosition)
}

func getMoveEndPosition(m chessboard.Move) realPosition {
	if m.SemiFirstHalf {
		return positionBetween(m.StartPosition, m.EndPosition)
	}
	return createRealPosition(m.EndPosition)
}
