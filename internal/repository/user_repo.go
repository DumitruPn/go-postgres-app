package repository

import (
	"database/sql"
	"go-postgres-app/internal/models"
)

func GetAllUsers(db *sql.DB) ([]models.User, error) {
	rows, err := db.Query(`SELECT id, first_name, last_name, age, car_id FROM data.users`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Age, &user.CarID); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func InsertUser(db *sql.DB, user models.User) (int, error) {
	var id int
	err := db.QueryRow(`
        INSERT INTO data.users (first_name, last_name, age, car_id)
        VALUES ($1, $2, $3, $4) RETURNING id
    `, user.FirstName, user.LastName, user.Age, user.CarID).Scan(&id)

	return id, err
}
