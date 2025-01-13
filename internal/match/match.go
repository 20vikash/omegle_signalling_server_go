package match

import (
	"errors"
	"math/rand"
	"signal/signal/internal/helper"
	"sync"

	"github.com/gorilla/websocket"
)

func Match_pair(conns *[]*websocket.Conn, mu *sync.Mutex) (*websocket.Conn, *websocket.Conn, error) {
	mu.Lock()

	if len(*conns) < 2 {
		mu.Unlock()
		return nil, nil, errors.New("not enough connections")
	}

	con1 := (*conns)[rand.Intn(len(*conns))]
	*conns = helper.RemoveSliceElement(*conns, con1)
	con2 := (*conns)[rand.Intn(len(*conns))]
	*conns = helper.RemoveSliceElement(*conns, con2)

	mu.Unlock()

	return con1, con2, nil
}

func Next_pair(conn *websocket.Conn) (*websocket.Conn, error) {

}
