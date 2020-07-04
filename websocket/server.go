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
		reversi, ok := games[data.GameId]
		if !ok {
			break
		}
		if reversi.Put(data.Turn, data.Point.X, data.Point.Y) {
			if err := websocket.JSON.Send(ws, reversi.Board); err != nil {
				log.Println(err)
			}
		}
	}
}
