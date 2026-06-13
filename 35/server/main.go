package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	defer conn.Close()

	log.Println("Client connected")

	go func() {
		for {
			time.Sleep(2 * time.Second)

			err := conn.WriteMessage(
				websocket.TextMessage,
				[]byte("Hello from server"),
			)
			if err != nil {
				return
			}
		}
	}()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Client disconnected")
			return
		}

		log.Printf("Client: %s", message)
	}
}

func main() {
	http.HandleFunc("/ws", wsHandler)

	log.Println("Server started on :8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
