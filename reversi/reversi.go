package reversi

type Reversi struct {
	Board
	Turn int
}

type PutResult string

const (
	NotYourTurn PutResult = "not_your_turn"
	InvalidPut  PutResult = "invalid_put"
	TurnChange  PutResult = "turn_change"
	TurnPass    PutResult = "turn_pass"
	GameEnd     PutResult = "game_end"
)

func Init() *Reversi {
	board := initBoard()
	board.suggest(1)
	rv := &Reversi{board, 1}
	rv.show()
	return rv
}

func (rv *Reversi) Put(t, x, y int) PutResult {
	if t != rv.Turn {
		return NotYourTurn
	}
	if !rv.Board.put(t, x, y) {
		return InvalidPut
	}
	if rv.Board.suggest(3 - rv.Turn) {
		rv.Turn = 3 - rv.Turn
		rv.show()
		return TurnChange
	} else if rv.Board.suggest(rv.Turn) {
		rv.show()
		return TurnPass
	} else {
		return GameEnd
	}
}

func (rv Reversi) BoardInfo() [][]int {
	return rv.Board.board()
}
