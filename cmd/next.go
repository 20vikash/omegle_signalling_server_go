package main

import (
	"log"
	"signal/signal/internal/match"

	"github.com/gorilla/websocket"
)

func Next_pair(ws *websocket.Conn) {
	on, tn, old, err := match.Next_pair(ws, &clients, &pairs, Mu)

	if err != nil {
		log.Println(err.Error())
		return
	}

	if tn == ws {
		pairs[tn] = on
		pairs[on] = tn

		writeNewPair(ws, on)

		return
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
