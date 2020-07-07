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
		if rv.Put(data.Turn, data.Point.X + 1, data.Point.Y + 1) == reversi.TurnChange {
			msg := SendMessage{
				Type:  "board",
				Data: rv.BoardInfo(),
			}
			if err := websocket.JSON.Send(ws, msg); err != nil {
				log.Println(err)
			}
		}
	}
}
