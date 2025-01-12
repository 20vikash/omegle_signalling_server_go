package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/ws", handleConnections)

	fmt.Println("WebSocket server started on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("ListenAndServe:", err)
	}
}
