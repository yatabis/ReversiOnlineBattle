package client

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"

	"ReversiOnlineBattle/websocket"
)

func Init(mux *http.ServeMux) {
	mux.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("./website/static/"))))
	mux.Handle("/waiting", http.HandlerFunc(waitingHandler))
	mux.Handle("/join", http.HandlerFunc(joinHandler))
	mux.Handle("/play", http.HandlerFunc(playHandler))
}

func waitingHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./website/static/html/waiting.html")
	if err != nil {
		log.Println(err)
		return
	}
	gameId := getCookie(r, "GameID")
	if gameId == "" {
		playerId := getCookie(r, "PlayerID")
		gameId, playerId = websocket.StartGame("", playerId)
		http.SetCookie(w, cookie("GameID", gameId))
		http.SetCookie(w, cookie("PlayerID", playerId))
	}
	if err := t.Execute(w, map[string]string{"GameID": gameId}); err != nil {
		log.Println(err)
	}
}

func joinHandler(w http.ResponseWriter, r *http.Request) {
	gameId := r.URL.Query().Get("gameid")
	if gameId == "" {
		http.ServeFile(w, r, "./website/static/html/join.html")
	} else {
		playerId := getCookie(r, "PlayerID")
		playerId = websocket.JoinGame(gameId, playerId)
		if playerId == "" {
			w.WriteHeader(http.StatusBadRequest)
			if _, err := fmt.Fprintf(w, "the game id `%s` is not started.\n", gameId); err != nil {
				log.Println(err)
			}
			return
		}
		http.SetCookie(w, cookie("GameID", gameId))
		http.SetCookie(w, cookie("PlayerID", playerId))
		w.WriteHeader(http.StatusOK)
		if _, err := fmt.Fprint(w, gameId); err != nil {
			log.Println(err)
		}
	}
}

func playHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./website/static/html/play.html")
}

func cookie(name, value string) *http.Cookie {
	return &http.Cookie{
		Name: name,
		Value: url.PathEscape(value),
		MaxAge: 60 * 60 * 24 * 14,
		SameSite: http.SameSiteStrictMode,
		// TODO: 本番デプロイするときは Secure 属性を true にする
	}
}

func getCookie(r *http.Request, name string) string {
	cookie, err := r.Cookie(name)
	if err != nil {
		return ""
	}
	value, err := url.PathUnescape(cookie.Value)
	if err != nil {
		return ""
	}
	return value
}
