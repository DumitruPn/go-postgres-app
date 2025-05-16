package handlers

import (
	"database/sql"
	"encoding/json"
	"go-postgres-app/internal/models"
	"go-postgres-app/internal/repository"
	"net/http"
)

func GetCars(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cars, err := repository.GetAllCars(db)
		if err != nil {
			http.Error(w, "Could not fetch cars", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(cars)
	}
}

func CreateCar(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var car models.Car
		if err := json.NewDecoder(r.Body).Decode(&car); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		id, err := repository.InsertCar(db, car)
		if err != nil {
			http.Error(w, "Could not insert car", http.StatusInternalServerError)
			return
		}

		car.ID = id
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(car)
	}
}
