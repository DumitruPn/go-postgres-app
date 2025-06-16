package ws

import (
	"golang.org/x/net/websocket"
	"log"
	"sync"
)

var (
	clients   = make(map[*websocket.Conn]bool)
	clientsMu sync.Mutex
)

func Handler(ws *websocket.Conn) {
	registerClient(ws)
	defer unregisterClient(ws)

	for {
		var msg string
		if err := websocket.Message.Receive(ws, &msg); err != nil {
			log.Println("Receive error:", err)
			break
		}

		log.Printf("Received: %s", msg)
		broadcast(msg)
	}
}

func registerClient(conn *websocket.Conn) {
	clientsMu.Lock()
	defer clientsMu.Unlock()
	clients[conn] = true
	log.Println("Client connected. Total clients:", len(clients))
}

func unregisterClient(conn *websocket.Conn) {
	clientsMu.Lock()
	defer clientsMu.Unlock()
	delete(clients, conn)
	log.Println("Client disconnected. Total clients:", len(clients))
}

func broadcast(message string) {
	clientsMu.Lock()
	defer clientsMu.Unlock()

	for conn := range clients {
		if err := websocket.Message.Send(conn, message); err != nil {
			log.Println("Send error:", err)
		}
	}
}
