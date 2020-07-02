package client

import (
	"html/template"
	"log"
	"net/http"
)

func Init(mux *http.ServeMux) {
	mux.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("./website/static/"))))
	mux.Handle("/play", http.HandlerFunc(playHandler))
}

func playHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Host: %s", r.Host)
	t, err := template.ParseFiles("./website/static/html/play.html")
	if err != nil {
		log.Println(err)
		return
	}
	if err := t.Execute(w, r.Host); err != nil {
		log.Println(err)
	}
}
