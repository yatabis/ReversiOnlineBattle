package websocket

import (
	"golang.org/x/net/websocket"
	"log"
	"net/http"

	"ReversiOnlineBattle/reversi"
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
	for {
		var data Data
		if err := websocket.JSON.Receive(ws, &data); err != nil {
			log.Printf("recieve data: %+v\n", data)
			log.Println(err)
			break
		}
		log.Printf("recieve: %+v\n", data)
		rv, ok := games[data.GameId]
		if !ok {
			break
		}
		result := rv.Put(data.Turn, data.Point.X+1, data.Point.Y+1)
		send(ws, "board", rv.BoardInfo())
		switch result {
		case reversi.NotYourTurn:
			send(ws, string(result), nil)
		case reversi.InvalidPut:
			send(ws, string(result), nil)
		case reversi.TurnChange:
			send(ws, string(result), rv.Turn)
		case reversi.TurnPass:
			send(ws, string(result), rv.Turn)
		case reversi.GameEnd:
			send(ws, string(result), nil)
		default:
			break
		}
	}
}

func send(ws *websocket.Conn, result string, data interface{}) {
	msg := SendMessage{
		Type: result,
		Data: data,
	}
	if err := websocket.JSON.Send(ws, msg); err != nil {
		log.Println(err)
	}
}
