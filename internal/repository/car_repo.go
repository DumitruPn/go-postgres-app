package repository

import (
	"database/sql"
	"go-postgres-app/internal/models"
)

func GetAllCars(db *sql.DB) ([]models.Car, error) {
	rows, err := db.Query("SELECT id, name, model, year FROM data.cars")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cars []models.Car
	for rows.Next() {
		var c models.Car
		if err := rows.Scan(&c.ID, &c.Name, &c.Model, &c.Year); err != nil {
			return nil, err
		}
		cars = append(cars, c)
	}
	return cars, nil
}

func InsertCar(db *sql.DB, car models.Car) (int, error) {
	var id int
	err := db.QueryRow("INSERT INTO data.cars (name, model, year) VALUES ($1, $2, $3) RETURNING id",
		car.Name, car.Model, car.Year).Scan(&id)
	return id, err
}
