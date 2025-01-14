package match

import (
	"errors"
	"math/rand"
	"signal/signal/internal/helper"
	"sync"

	"github.com/gorilla/websocket"
)

func Match_pair(conns *[]*websocket.Conn, pairs *map[*websocket.Conn]*websocket.Conn, mu *sync.Mutex) (*websocket.Conn, *websocket.Conn, error) {
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

func Next_pair(conn *websocket.Conn, conns *[]*websocket.Conn, pairs *map[*websocket.Conn]*websocket.Conn, mu *sync.Mutex) (*websocket.Conn, *websocket.Conn, *websocket.Conn, error) {
	mu.Lock()
	defer mu.Unlock()

	conn2, exists := (*pairs)[conn]

	if !exists {
		return nil, nil, conn2, errors.New("no connections found yet")
	}

	delete((*pairs), conn)
	delete((*pairs), conn2)

	if len(*conns) == 0 {
		return conn2, conn, conn2, nil
	}

	con1 := (*conns)[rand.Intn(len(*conns))]
	*conns = helper.RemoveSliceElement(*conns, con1)

	if len(*conns) == 0 {
		*conns = append(*conns, conn2)
		return con1, nil, conn2, nil
	}

	con2 := (*conns)[rand.Intn(len(*conns))]
	*conns = helper.RemoveSliceElement(*conns, con2)

	return con1, con2, conn2, nil
}
