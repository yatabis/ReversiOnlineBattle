package client

import (
	"html/template"
	"log"
	"net/http"

	"ReversiOnlineBattle/websocket"
)

func Init(mux *http.ServeMux) {
	mux.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("./website/static/"))))
	mux.Handle("/waiting", http.HandlerFunc(waitingHandler))
	mux.Handle("/play", http.HandlerFunc(playHandler))
}

func waitingHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./website/static/html/waiting.html")
	if err != nil {
		log.Println(err)
		return
	}
	gameId := ""
	if g, err := r.Cookie("GameID"); err == nil {
		gameId = g.Value
	} else {
		playerId := ""
		if p, err := r.Cookie("PlayerID"); err == nil {
			playerId = p.Value
		}
		gameId, playerId = websocket.StartGame("", playerId)
		http.SetCookie(w, cookie("GameID", gameId))
		http.SetCookie(w, cookie("PlayerID", playerId))
	}
	if err := t.Execute(w, map[string]string{"GameID": gameId}); err != nil {
		log.Println(err)
	}
}

func playHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./website/static/html/play.html")
}

func cookie(name, value string) *http.Cookie {
	return &http.Cookie{
		Name: name,
		Value: value,
		MaxAge: 60 * 60 * 24 * 14,
		SameSite: http.SameSiteStrictMode,
		// TODO: 本番デプロイするときは Secure 属性を true にする
	}
}
