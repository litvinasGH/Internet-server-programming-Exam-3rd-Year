package main

import (
	"log"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	conn, _, err := websocket.DefaultDialer.Dial(
		"ws://localhost:8080/ws",
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	go func() {
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				return
			}

			log.Printf("Server: %s", message)
		}
	}()

	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()

	timeout := time.After(10 * time.Second)

	for {
		select {

		case <-ticker.C:
			err := conn.WriteMessage(
				websocket.TextMessage,
				[]byte("Hello from client"),
			)
			if err != nil {
				return
			}

		case <-timeout:
			log.Println("Client closing connection")
			return
		}
	}
}
