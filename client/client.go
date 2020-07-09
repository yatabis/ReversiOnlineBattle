package client

import (
	"html/template"
	"log"
	"net/http"

	"ReversiOnlineBattle/websocket"
)

func Init(mux *http.ServeMux) {
	mux.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("./website/static/"))))
	mux.Handle("/play", http.HandlerFunc(playHandler))
}

func playHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./website/static/html/play.html")
	if err != nil {
		log.Println(err)
		return
	}
	if _, err := r.Cookie("gameID"); err != nil {
		gameId := websocket.StartGame("")
		http.SetCookie(w, &http.Cookie{
			Name: "gameID",
			Value: gameId,
			MaxAge: 60 * 60 * 24 * 14,
		})
	}
	if err := t.Execute(w, nil); err != nil {
		log.Println(err)
	}
}
