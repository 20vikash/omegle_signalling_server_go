package helper

import "github.com/gorilla/websocket"

func RemoveSliceElement(s []*websocket.Conn, element any) []*websocket.Conn {
	res := make([]*websocket.Conn, 0)

	for _, v := range s {
		if v == element {
			continue
		}

		res = append(res, v)
	}

	return res
}
