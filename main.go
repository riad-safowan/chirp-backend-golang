package main

import (
	"fmt"
	"net/http"

	"github.com/riad-safowan/chirp-backend/pkg/websocket"
)

func main() {
	setupRoutes()
	http.ListenAndServe(":9090", nil)
}

func setupRoutes() {
	pool := websocket.NewPool()
	go pool.Start()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWS(pool, w, r)
	})
}

func serveWS(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	conn, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Print("cannot serve")
		return 
	}

	client:= &websocket.Client{
		Conn: conn,
		Pool: pool,
	}

	pool.Register <- client
	client.Read()
}
