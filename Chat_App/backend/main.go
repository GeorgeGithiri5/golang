package main

import (
	"fmt"
	"net/http"

	"github.com/pkg/websocket"
)

func serveWs(w http.ResponseWriter, r *http.Request) {
	ws, err := websocket.Upgrader(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+V\n", err)
	}
	go websocket.Writer(ws)
	websocket.Reader(ws)
}

func setupRoutes() {
	http.HandleFunc("/ws", serveWs)
}

func main() {
	fmt.Println("Distributed Chat App")
	setupRoutes()
	http.ListenAndServe(":8080", nil)
}
