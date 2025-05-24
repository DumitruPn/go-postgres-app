package user

import (
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"go-postgres-app/internal/car"
	"net/http"
)

type Handler struct {
	db *sql.DB
}

func NewHandler(db *sql.DB) *Handler {
	return &Handler{db}
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var input CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if input.FirstName == "" || input.LastName == "" || input.Age == 0 {
		http.Error(w, "Missing user data", http.StatusBadRequest)
		return
	}

	id, err := Insert(h.db, input)
	if err != nil {
		http.Error(w, "Failed to insert user", http.StatusInternalServerError)
		return
	}

	user := Dto{
		ID:        id,
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Age:       input.Age,
		Cars:      []car.Dto{},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	users, err := GetAll(h.db)
	if err != nil {
		http.Error(w, "Failed to fetch users", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(AsDtos(users))
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, ok := vars["id"]
	if !ok {
		http.Error(w, "Missing user id", http.StatusBadRequest)
		return
	}

	u, err := Get(h.db, userId)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "User not found", http.StatusNotFound)
		} else {
			http.Error(w, "Failed to fetch user", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(AsDto(*u))
}
