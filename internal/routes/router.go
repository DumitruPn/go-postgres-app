package routes

import (
	"database/sql"
	"go-postgres-app/internal/handlers"
	"net/http"
)

func Setup(db *sql.DB) *http.ServeMux {
	mux := http.NewServeMux()

	// Register cars routes
	registerCarRoutes(mux, db)

	// Register users routes
	registerUserRoutes(mux, db)

	return mux
}

func registerCarRoutes(mux *http.ServeMux, db *sql.DB) {
	mux.HandleFunc("/cars", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.GetCars(db)(w, r)
		case http.MethodPost:
			handlers.CreateCar(db)(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})
}

func registerUserRoutes(mux *http.ServeMux, db *sql.DB) {
	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.GetUsers(db)(w, r)
		case http.MethodPost:
			handlers.CreateUser(db)(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})
}
