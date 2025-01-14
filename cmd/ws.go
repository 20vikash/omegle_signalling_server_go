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
var pairs = make(map[*websocket.Conn]*websocket.Conn)

func handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println(err)
		return
	}

	clients = append(clients, ws)

	log.Println(clients)

	con1, con2, err := match.Match_pair(&clients, &pairs, Mu)
	pairs[con1] = con2
	pairs[con2] = con1

	if err != nil {
		log.Println("No pair found.. Waiting for another connection")
	}

	if con1 != nil && con2 != nil {
		con1.WriteMessage(websocket.TextMessage, []byte("You got a pair"))
		con2.WriteMessage(websocket.TextMessage, []byte("You got a pair"))
	}

	for {
		_, p, err := ws.ReadMessage()
		if err != nil {
			log.Println(err)
		}

		message := string(p)

		if message == "NEXT" {
			on, tn, old, err := match.Next_pair(ws, &clients, &pairs, Mu)

			if err != nil {
				log.Println(err.Error())
				continue
			}

			if tn == ws {
				pairs[tn] = on
				pairs[on] = tn

				writeNewPair(ws, on)

				continue
			}

			pairs[ws] = on
			pairs[on] = ws
			writeNewPair(ws, on)

			if tn != nil {
				pairs[old] = tn
				pairs[tn] = old
				writeNewPair(old, tn)
			}
		}
	}
}

func writeNewPair(con1 *websocket.Conn, con2 *websocket.Conn) {
	con1.WriteMessage(websocket.TextMessage, []byte("You got a new pair"))
	con2.WriteMessage(websocket.TextMessage, []byte("You got a new pair"))
}
