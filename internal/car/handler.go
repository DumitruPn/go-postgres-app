package car

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

func GetCarsHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cars, err := GetAll(db)
		if err != nil {
			http.Error(w, "Could not fetch cars", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(cars)
	}
}

func CreateCarHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var car Car
		if err := json.NewDecoder(r.Body).Decode(&car); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		id, err := Insert(db, car)
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
