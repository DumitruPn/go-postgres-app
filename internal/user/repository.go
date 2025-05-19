package user

import (
	"database/sql"
)

func GetAll(db *sql.DB) ([]User, error) {
	rows, err := db.Query(`SELECT id, first_name, last_name, age, car_id FROM data.users`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Age, &user.CarID); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func Insert(db *sql.DB, user User) (int, error) {
	var id int
	err := db.QueryRow(`
        INSERT INTO data.users (first_name, last_name, age, car_id)
        VALUES ($1, $2, $3, $4) RETURNING id
    `, user.FirstName, user.LastName, user.Age, user.CarID).Scan(&id)

	return id, err
}
