package main

import (
	"log"
	"net/http"
	"signal/signal/internal/helper"
	"signal/signal/internal/match"
	"strings"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

var clients = make([]*websocket.Conn, 0)
var pairs = make(map[*websocket.Conn]*websocket.Conn)

func handleConnections(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "https://rtc.selfmade.fun")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	ws, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println(err)
		return
	}

	clients = append(clients, ws)

	log.Println(clients)

	con1, con2, err := match.Match_pair(&clients, &pairs, Mu)

	if err != nil {
		log.Println("No pair found.. Waiting for another connection")
	}

	if con1 != nil && con2 != nil {
		pairs[con1] = con2
		pairs[con2] = con1
		Initiate(con1)
		log.Println("Step 1 done")
	}

	for {
		_, p, err := ws.ReadMessage()
		if err != nil {
			log.Println(err)
		}

		message := string(p)

		if message == "NEXT" {
			Next_pair(ws)
		} else {
			pair := strings.Split(message, "~")
			role := pair[0]
			code := pair[1]

			if role == helper.OFFER {
				SDP_offer(pairs[ws], code)
				log.Println("Step 2 done")
			} else if role == helper.ANSWER {
				log.Println("Hey there")
				SDP_answer(pairs[ws], code)
				log.Println("Step 3 done")
			}
		}
	}
}
