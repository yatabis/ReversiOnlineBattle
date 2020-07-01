package client

import (
	"net/http"
)

func Init(mux *http.ServeMux) {
	mux.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("./website/static/"))))
	mux.Handle("/play", http.HandlerFunc(playHandler))
}

func playHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./website/static/html/play.html")
}
