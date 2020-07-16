package websocket

import (
	"golang.org/x/net/websocket"
	"log"

	"ReversiOnlineBattle/reversi"
)

type GameStatus string

const (
	StatusWaiting  GameStatus = "waiting"
	StatusStarting GameStatus = "starting"
	StatusPlaying  GameStatus = "playing"
	StatusClosed   GameStatus = "closed"
)

type GameRoom struct {
	GameId    string
	HostId    string
	GuestId   string
	Status    GameStatus
	HostConn  *websocket.Conn
	GuestConn *websocket.Conn
	*reversi.Reversi
}

var games = make(map[string]*GameRoom)

var players = make(map[string]string)

func StartGame(gameId, hostId string) (string, string) {
	if gameId == "" {
		gameId = createGameId()
	}
	if hostId == "" {
		hostId = createPlayerId()
	}
	log.Printf("Host ID: %s\n", hostId)
	games[gameId] = &GameRoom{
		GameId:  gameId,
		HostId:  hostId,
		Status:  StatusWaiting,
		Reversi: reversi.Init(),
	}
	players[hostId] = gameId
	return gameId, hostId
}

func JoinGame(gameId, guestId string) string {
	game, ok := games[gameId]
	if !ok {
		return ""
	}
	if guestId == "" {
		guestId = createPlayerId()
	}
	game.GuestId = guestId
	players[guestId] = gameId
	game.Status = StatusStarting
	return guestId
}

func (g *GameRoom) onMessage(ws *websocket.Conn) {
	for {
		var point Point
		if err := websocket.JSON.Receive(ws, &point); err != nil {
			log.Printf("recieve point: %+v\n", point)
			log.Println(err)
			break
		}
		log.Printf("recieve: %+v\n", point)
		turn := 0
		switch ws {
		case g.HostConn:
			turn = 1
		case g.GuestConn:
			turn = 2
		default:
			log.Printf("receive a message from invalid connection: %+v\n", ws)
		}
		result := g.Reversi.Put(turn, point.X+1, point.Y+1)
		switch result {
		case reversi.NotYourTurn:
			g.send(ws, string(result), nil)
		case reversi.InvalidPut:
			g.send(ws, string(result), nil)
		case reversi.TurnChange:
			g.sendBoth("board", g.Reversi.BoardInfo())
			g.sendBoth(string(result), g.Reversi.Turn)
		case reversi.TurnPass:
			g.sendBoth("board", g.Reversi.BoardInfo())
			g.sendBoth(string(result), g.Reversi.Turn)
		case reversi.GameEnd:
			g.sendBoth("board", g.Reversi.BoardInfo())
			g.sendBoth(string(result), nil)
		default:
			break
		}
	}
}

func (g *GameRoom) send(ws *websocket.Conn, t string, data interface{}) {
	msg := SendMessage{
		Type: t,
		Data: data,
	}
	if err := websocket.JSON.Send(ws, msg); err != nil {
		log.Println(err)
	}
}

func (g *GameRoom) sendBoth(t string, data interface{}) {
	g.send(g.HostConn, t, data)
	g.send(g.GuestConn, t, data)
}
