package main

import (
	"log"
	"net/http"
	"signal/signal/internal/match"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

var clients = make([]*websocket.Conn, 0)

func handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println(err)
		return
	}

	clients = append(clients, ws)

	log.Println(clients)

	con1, con2, err := match.Match_pair(&clients, Mu)
	if err != nil {
		log.Println("No pair found.. Waiting for another connection")
	}

	if con1 != nil && con2 != nil {
		con1.WriteMessage(websocket.TextMessage, []byte("You got a pair"))
		con2.WriteMessage(websocket.TextMessage, []byte("You got a pair"))
	}
}
