package main

import (
	"fmt"
	"signal/signal/internal/helper"

	"github.com/gorilla/websocket"
)

func Initiate(con1 *websocket.Conn) {
	con1.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("%v", helper.INITIATE)))
}

func SDP_offer(con2 *websocket.Conn, code string) {

}

func SDP_answer(con1 *websocket.Conn, code string) {

}
