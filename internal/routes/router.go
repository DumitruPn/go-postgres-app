package routes

import (
	"database/sql"
	"go-postgres-app/internal/car"
	"go-postgres-app/internal/user"
	"net/http"
)

func Setup(db *sql.DB) *http.ServeMux {
	mux := http.NewServeMux()

	registerCarRoutes(mux, db)

	registerUserRoutes(mux, db)

	return mux
}

func registerCarRoutes(mux *http.ServeMux, db *sql.DB) {
	mux.HandleFunc("/cars", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			car.GetCarsHandler(db)(w, r)
		case http.MethodPost:
			car.CreateCarHandler(db)(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})
}

func registerUserRoutes(mux *http.ServeMux, db *sql.DB) {
	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			user.GetUsersHandler(db)(w, r)
		case http.MethodPost:
			user.CreateUserHandler(db)(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})
}
