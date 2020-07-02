package websocket

import (
	"golang.org/x/net/websocket"
	"io"
	"log"
	"net/http"
)

func Init(mux *http.ServeMux) {
	mux.Handle("/open", websocket.Handler(open))
}

func open(conn *websocket.Conn) {
	n, err := io.Copy(conn, conn)
	log.Println(n)
	if err != nil {
		log.Println(err)
	}
}
