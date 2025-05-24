package user

import (
	"database/sql"
)

func GetAll(db *sql.DB) ([]User, error) {
	rows, err := db.Query(`SELECT id, first_name, last_name, age FROM data.users`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Age); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func Get(db *sql.DB, id string) (*User, error) {
	row := db.QueryRow(`SELECT id, first_name, last_name, age FROM data.users where id = $1`, id)

	var user User
	if err := row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Age); err != nil {
		return nil, err
	}

	return &user, nil
}

func Insert(db *sql.DB, user CreateUserRequest) (int, error) {
	var id int
	err := db.QueryRow(`
        INSERT INTO data.users (first_name, last_name, age)
        VALUES ($1, $2, $3) RETURNING id
    `, user.FirstName, user.LastName, user.Age).Scan(&id)

	return id, err
}
