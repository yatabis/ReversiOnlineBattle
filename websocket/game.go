package websocket

import "ReversiOnlineBattle/reversi"

var games = make(map[string]*reversi.Reversi)

func StartGame() string {
	gameId := "test"
	games[gameId] = reversi.Init()
	return gameId
}
