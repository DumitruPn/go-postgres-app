package user

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

func GetUsersHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := GetAll(db)
		if err != nil {
			http.Error(w, "Failed to fetch users", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)
	}
}

func CreateUserHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}

		if user.FirstName == "" || user.LastName == "" || user.Age == 0 || user.CarID == 0 {
			http.Error(w, "Missing user data", http.StatusBadRequest)
			return
		}

		id, err := Insert(db, user)
		if err != nil {
			http.Error(w, "Failed to insert user", http.StatusInternalServerError)
			return
		}

		user.ID = id
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(user)
	}
}
