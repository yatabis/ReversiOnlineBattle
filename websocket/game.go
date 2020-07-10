package websocket

import (
	"golang.org/x/net/websocket"
	"log"

	"ReversiOnlineBattle/reversi"
)


type GameRoom struct {
	GameId   string
	HostId   string
	ClientId string
	ws       *websocket.Conn
	*reversi.Reversi
}

var games = make(map[string]*GameRoom)

var players = make(map[string]string)

func createGameId() string {
	return "test"
}

func createPlayerId() string {
	return "player"
}

func StartGame(gameId, playerId string) (string, string) {
	if gameId == "" {
		gameId = createGameId()
	}
	if playerId == "" {
		playerId = createPlayerId()
	}
	games[gameId] = &GameRoom{
		GameId:   gameId,
		HostId:   playerId,
		Reversi: reversi.Init(),
	}
	players[playerId] = gameId
	return gameId, playerId
}

func (g *GameRoom) onMessage() {
	for {
		var data Data
		if err := websocket.JSON.Receive(g.ws, &data); err != nil {
			log.Printf("recieve data: %+v\n", data)
			log.Println(err)
			break
		}
		log.Printf("recieve: %+v\n", data)
		result := g.Reversi.Put(data.Turn, data.Point.X+1, data.Point.Y+1)
		g.send("board", g.Reversi.BoardInfo())
		switch result {
		case reversi.NotYourTurn:
			g.send(string(result), nil)
		case reversi.InvalidPut:
			g.send(string(result), nil)
		case reversi.TurnChange:
			g.send(string(result), g.Reversi.Turn)
		case reversi.TurnPass:
			g.send(string(result), g.Reversi.Turn)
		case reversi.GameEnd:
			g.send(string(result), nil)
		default:
			break
		}
	}
}

func (g *GameRoom) send(t string, data interface{}) {
	msg := SendMessage{
		Type: t,
		Data: data,
	}
	if err := websocket.JSON.Send(g.ws, msg); err != nil {
		log.Println(err)
	}
}
