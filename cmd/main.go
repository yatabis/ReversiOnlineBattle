package main

import (
	"fmt"
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
		_, _ = fmt.Fprintf(w, "Hello, this is Reversi Online Battle.\nplay game => /play\n")
	})
	websocket.Init(mux)
	client.Init(mux)
	log.Fatal(http.ListenAndServe(":" + port, mux))
}
