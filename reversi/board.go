package reversi

import "fmt"

type Board [8][8]int

func initBoard() Board {
	rv := Board{}
	rv[3][3] = 2
	rv[3][4] = 1
	rv[4][3] = 1
	rv[4][4] = 2
	rv.show()
	return rv
}

func (b Board) show() {
	for _, row := range b {
		fmt.Println(row)
	}
}

func (b *Board) put(t, x, y int) bool {
	if b[y][x] != 0 {
		return false
	}
	if !b.reverse(t, x, y) {
		return false
	}
	b[y][x] = t
	//fmt.Printf("(%d, %d) <- %d\n", x, y, t)
	b.show()
	return true
}

func (b *Board) reverse(t, x, y int) bool {
	c := 0
	for dy := -1; dy <= 1; dy += 1 {
		for dx := -1; dx <= 1; dx += 1 {
			if dx == 0 && dy == 0 {
				continue
			}
			n := b.search(t, x + dx, y + dy, dx, dy)
			if n > 1 {
				c += n - 1
			}
		}
	}
	//fmt.Printf("c = %d\n", c)
	return c > 0
}

func (b *Board) search(t, x, y, dx, dy int) int {
	//fmt.Printf("searching (%d, %d)\n", x, y)
	if x + dx < 0 || x + dx > 7 || y + dy < 0 || y + dy > 7 {
		return 0
	}
	switch b[y][x] {
	case 0:
		return 0
	case t:
		return 1
	case 3 - t:
		n := b.search(t, x + dx, y + dy, dx, dy)
		if n == 0 {
			return 0
		}
		b[y][x] = t
		//fmt.Printf("(%d, %d) <- %d\n", x, y, t)
		return n + 1
	default:
		return 0
	}
}
