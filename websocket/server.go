package websocket

import (
	"golang.org/x/net/websocket"
	"log"
	"net/http"
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
	mux.Handle("/open", websocket.Handler(open))
}

func open(ws *websocket.Conn) {
	var gameId string
	if err := websocket.Message.Receive(ws, &gameId); err != nil {
		log.Printf("failed to receive the game id: %e\n", err)
	}
	log.Printf("connected to %s\n", gameId)
	game, ok := games[gameId]
	if !ok {
		StartGame(gameId)
		game = games[gameId]
	}
	game.ws = ws
	game.onMessage()
}
