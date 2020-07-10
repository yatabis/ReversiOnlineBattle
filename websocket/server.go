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
	time.Sleep(5 * time.Second)
	if err := websocket.Message.Send(ws, "start"); err != nil {
		log.Printf("failed to send start-message: %e\n", err)
	}
}

func open(ws *websocket.Conn) {
	var playerId string
	if err := websocket.Message.Receive(ws, &playerId); err != nil {
		log.Printf("failed to receive the player id: %e\n", err)
	}
	log.Printf("connected to %s\n", playerId)
	gameId, okPlayer := players[playerId]
	game, okGame := games[gameId]
	if !(okPlayer && okGame) {
		gameId, _ = StartGame(gameId, playerId)
		game = games[gameId]
	}
	log.Printf("loaded game %s\n", gameId)
	game.ws = ws
	game.send("board", game.Reversi.BoardInfo())
	game.onMessage()
}
