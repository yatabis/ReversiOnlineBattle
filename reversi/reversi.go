package reversi

type Reversi struct {
	Board
	turn int
}

type PutResult string

const (
	InvalidPut PutResult = "invalid_put"
	TurnChange PutResult = "turn_change"
)

func Init() *Reversi {
	board := initBoard()
	board.suggest(1)
	rv := &Reversi{board, 1}
	rv.show()
	return rv
}

func (rv *Reversi) Put(t, x, y int) PutResult {
	if t != rv.turn {
		return InvalidPut
	}
	if !rv.Board.put(t, x, y) {
		return InvalidPut
	}
	rv.turn = 3 - rv.turn
	rv.Board.suggest(rv.turn)
	rv.show()
	return TurnChange
}

func (rv Reversi) BoardInfo() [][]int {
	return rv.Board.board()
}
