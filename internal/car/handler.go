package car

import (
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"go-postgres-app/internal/users_cars"
	"net/http"
)

type Handler struct {
	db *sql.DB
}

func NewHandler(db *sql.DB) *Handler {
	return &Handler{db}
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var input CreateCarRequest
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if input.Name == "" || input.Model == "" || input.Year == 0 {
		http.Error(w, "Missing car data", http.StatusBadRequest)
		return
	}

	id, err := Insert(h.db, input)
	if err != nil {
		http.Error(w, "Failed to insert car", http.StatusInternalServerError)
		return
	}

	car := Dto{
		ID:    id,
		Name:  input.Name,
		Model: input.Model,
		Year:  input.Year,
		Users: []users_cars.DtoUserWithoutCar{},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(car)
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	carId, ok := vars["id"]
	if !ok {
		http.Error(w, "Missing car id", http.StatusBadRequest)
	}

	u, err := Get(h.db, carId)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "Car not found", http.StatusNotFound)
		} else {
			http.Error(w, "Failed to fetch car", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(AsDtoWithUsers(*u))
}

func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	cars, err := GetAll(h.db)
	if err != nil {
		http.Error(w, "Failed to fetch cars", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(AsDtosWithUsers(cars))
}
