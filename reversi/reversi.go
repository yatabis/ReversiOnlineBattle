package reversi

type Reversi struct {
	Board
	turn int
}

func Init() *Reversi {
	board := initBoard()
	board.suggest(1)
	rv := &Reversi{board, 1}
	rv.show()
	return rv
}

func (rv *Reversi) Put(t, x, y int) bool {
	if t != rv.turn {
		return false
	}
	if !rv.Board.put(t, x, y) {
		return false
	}
	rv.turn = 3 - rv.turn
	rv.Board.suggest(rv.turn)
	rv.show()
	return true
}

func (rv Reversi) BoardInfo() [][]int {
	return rv.Board.board()
}
