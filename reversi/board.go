package reversi

import (
	"fmt"
)

type Board [10][10]int

func initBoard() Board {
	b := Board{}
	for i := 0; i < 10; i++ {
		b[0][i] = -1
		b[9][i] = -1
		b[i][0] = -1
		b[i][9] = -1
	}
	b[4][4] = 2
	b[4][5] = 1
	b[5][4] = 1
	b[5][5] = 2
	b.show()
	return b
}

func (b Board) board() (board [][]int) {
	board = make([][]int, 8)
	for i := 0; i < 8; i++ {
		board[i] = b[i + 1][1:9]
	}
	return
}

func (b Board) show() {
	for _, row := range b.board() {
		fmt.Println(row)
	}
}

func (b *Board) put(t, x, y int) bool {
	if b[y][x] != 0 {
		return false
	}
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if dx == 0 && dy == 0 {
				continue
			}
			n := b.search(t, x, y, dx, dy)
			if n == 0 {
				continue
			}
			b.reverse(t, x, y, dx, dy, n + 1)
		}
	}
	b.show()
	return b[y][x] == t
}

func (b *Board) search(t, x, y, dx, dy int) int {
	//log.Printf("searching for (%d, %d) for (%d, %d)...\n", x, y, dx, dy)
	n := 0
	for b[y + dy][x + dx] == 3 - t {
		x += dx
		y += dy
		n += 1
		//log.Printf("  b[%d][%d] = %d (n = %d)", x, y, b[y][x], n)
	}
	if b[y + dy][x + dx] == t {
		//log.Printf("  b[%d][%d] = %d (n = %d): returned", x + dx, y + dy, b[y + dy][x + dx], n)
		return n
	} else {
		//log.Printf("  b[%d][%d] = %d (n = %d): returned", x + dx, y + dy, b[y + dy][x + dx], 0)
		return 0
	}
}

func (b *Board) reverse(t, x, y, dx, dy, n int) {
	for i := 0; i < n; i++ {
		b[y + dy * i][x + dx * i] = t
	}
}
