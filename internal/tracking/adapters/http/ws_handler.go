package http

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WsHandler(w http.ResponseWriter, r *http.Request) {
	followerIDStr := r.URL.Query().Get("followerId")
	followerID, err := strconv.Atoi(followerIDStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("Invalid id"))
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("Failed to upgrade connection"))
		return
	}

	// wsConn := &ws.WsConn{
	// 	FollowerID: followerID,
	// 	Conn: conn,
	// }

	go func() {
		defer conn.Close()
		for {
			var msg map[string]any
			if err := conn.ReadJSON(&msg); err != nil {
				break
			}
			fmt.Println("Client:", followerID, "| Received =>", msg)
		}
	}()
}
