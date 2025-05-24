package user

import (
	"database/sql"
	"go-postgres-app/internal/car"
	"maps"
	"slices"
)

func GetAll(db *sql.DB) ([]User, error) {
	rows, err := db.Query(`
		SELECT u.id, u.first_name, u.last_name, u.age, c.id, c.name, c.model, c.year FROM data.users u
		INNER JOIN data.users_cars uc ON u.id = uc.user_id
	  	INNER JOIN data.cars c ON uc.car_id = c.id
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := map[int]User{}
	for rows.Next() {
		var u User
		var c car.Car
		if err := rows.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Age, &c.ID, &c.Name, &c.Model, &c.Year); err != nil {
			return nil, err
		}
		existing, ok := users[u.ID]
		if ok {
			existing.Cars = append(existing.Cars, c)
			users[u.ID] = existing
		} else {
			u.Cars = []car.Car{c}
			users[u.ID] = u
		}
	}

	return slices.Collect(maps.Values(users)), nil
}

func Get(db *sql.DB, id string) (*User, error) {
	rows, err := db.Query(`
		SELECT u.id, u.first_name, u.last_name, u.age, c.id, c.name, c.model, c.year FROM data.users u
		INNER JOIN data.users_cars uc ON u.id = uc.user_id
	  	INNER JOIN data.cars c ON uc.car_id = c.id
		WHERE u.id = $1`,
		id,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var u User
	var cars []car.Car
	for rows.Next() {
		var c car.Car
		if err := rows.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Age, &c.ID, &c.Name, &c.Model, &c.Year); err != nil {
			return nil, err
		}
		cars = append(cars, c)
	}
	u.Cars = cars

	return &u, nil
}

func Insert(db *sql.DB, user CreateUserRequest) (int, error) {
	var id int
	err := db.QueryRow(`
        INSERT INTO data.users (first_name, last_name, age)
        VALUES ($1, $2, $3) RETURNING id
    `, user.FirstName, user.LastName, user.Age).Scan(&id)

	return id, err
}
