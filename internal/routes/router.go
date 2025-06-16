package routes

import (
	"database/sql"
	"github.com/gorilla/mux"
	"go-postgres-app/internal/car"
	"go-postgres-app/internal/user"
	"go-postgres-app/internal/ws"
	"golang.org/x/net/websocket"
	"net/http"
)

func Setup(db *sql.DB) *mux.Router {
	router := mux.NewRouter()

	// websockets
	router.Handle("/ws", websocket.Handler(ws.Handler))

	// user	endpoints
	userHandler := user.NewHandler(db)
	router.HandleFunc("/users", userHandler.Create).Methods(http.MethodPost)
	router.HandleFunc("/users/{id}", userHandler.Get).Methods(http.MethodGet)
	router.HandleFunc("/users", userHandler.GetAll).Methods(http.MethodGet)

	// car endpoints
	carHandler := car.NewHandler(db)
	router.HandleFunc("/cars", carHandler.Create).Methods(http.MethodPost)
	router.HandleFunc("/cars/{id}", carHandler.Get).Methods(http.MethodGet)
	router.HandleFunc("/cars", carHandler.GetAll).Methods(http.MethodGet)

	return router
}
