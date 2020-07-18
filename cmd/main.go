package main

import (
	"log"
	"net/http"
	"os"

	"ReversiOnlineBattle/client"
	"ReversiOnlineBattle/websocket"
)

func main() {
	port := os.Getenv("PORT")
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./website/static/html/top.html")
	})
	websocket.Init(mux)
	client.Init(mux)
	log.Fatal(http.ListenAndServe(":" + port, mux))
}
