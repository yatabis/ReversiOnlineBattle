package websocket

import (
	"golang.org/x/net/websocket"
	"log"
	"net/http"
	"time"
)

type Data struct {
	GameId string `json:"gameId"`
	Turn   int    `json:"turn"`
	Point  struct{
		X int `json:"x"`
		Y int `json:"Y"`
	}             `json:"point"`
}

type SendMessage struct {
	Type string      `json:"type"`
	Data interface{} `json:"data,omitempty"`
}

func Init(mux *http.ServeMux) {
	mux.Handle("/wait", websocket.Handler(wait))
	mux.Handle("/open", websocket.Handler(open))
}

func wait(ws *websocket.Conn) {
	var playerId string
	if err := websocket.Message.Receive(ws, &playerId); err != nil {
		log.Printf("failed to receive the player id: %e\n", err)
	}
	gameId, okPlayer := players[playerId]
	game, okGame := games[gameId]
	if !(okPlayer && okGame) {
		gameId, _ = StartGame(gameId, playerId)
		game = games[gameId]
	}
	for {
		if game.Status == StatusStarting {
			if err := websocket.Message.Send(ws, "start"); err != nil {
				log.Printf("failed to send start-message: %e\n", err)
			}
			game.Status = StatusPlaying
			break
		}
		time.Sleep(10 * time.Millisecond)
	}
}

func open(ws *websocket.Conn) {
	var gameId string
	if err := websocket.Message.Receive(ws, &gameId); err != nil {
		log.Printf("failed to receive the game id: %e\n", err)
	}
	game := games[gameId]
	game.ws = ws
	game.send("board", game.Reversi.BoardInfo())
	game.onMessage()
}
