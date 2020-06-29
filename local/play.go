package local

import (
	"fmt"
	"strconv"

	"ReversiOnlineBattle/reversi"
)

func Play() {
	rv := reversi.Init()
	t := 1
	for {
		var s string
		if t == 1 {
			fmt.Printf("黒の番：")
		} else {
			fmt.Printf("白の番：")
		}
		n, err := fmt.Scanf("%s", &s)
		if n != 1 || err != nil {
			break
		}
		if s == "q" {
			break
		}
		x, err := strconv.Atoi(s[0:1])
		if err != nil {
			fmt.Println("不正な入力です。")
			continue
		}
		y := int(s[1] - 97)
		if y < 0 || y > 7 {
			fmt.Println("不正な入力です。")
			continue
		}
		if !rv.Put(t, x - 1, y) {
			fmt.Println("不正な入力です。")
			continue
		}
		t = 3 - t
	}
	fmt.Println("終了します。")
}
