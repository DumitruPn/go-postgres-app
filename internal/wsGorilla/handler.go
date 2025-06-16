package wsGorilla

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/websocket"
	"go-postgres-app/internal/notification"
	"log"
	"net/http"
	"sync"
)

type Handler struct {
	db *sql.DB
}

var (
	clients   = make(map[*websocket.Conn]bool)
	clientsMu sync.Mutex
	upgrader  = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func NewHandler(db *sql.DB) *Handler {
	return &Handler{db}
}

func (h *Handler) Handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket upgrade error:", err)
		return
	}

	registerClient(conn)
	defer unregisterClient(conn)

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			break
		}

		input := notification.CreateNotificationRequest{Value: string(msg)}

		id, err := notification.Insert(h.db, input)
		if err != nil {
			http.Error(w, "Failed to insert notification", http.StatusInternalServerError)
			return
		}

		dto := notification.Dto{
			ID:    id,
			Value: input.Value,
		}
		jsonMsg, _ := json.Marshal(dto)

		log.Printf("Received: %s", jsonMsg)
		broadcast(jsonMsg)
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
	conn.Close()
	log.Println("Client disconnected. Total clients:", len(clients))
}

func broadcast(message []byte) {
	clientsMu.Lock()
	defer clientsMu.Unlock()

	for conn := range clients {
		if err := conn.WriteMessage(websocket.TextMessage, message); err != nil {
			log.Println("Broadcast error:", err)
		}
	}
}
