package user

import (
	"database/sql"
	"go-postgres-app/internal/users_cars"
	"maps"
	"slices"
)

func Insert(db *sql.DB, user CreateUserRequest) (int, error) {
	var id int
	err := db.QueryRow(`
        INSERT INTO data.users (first_name, last_name, age)
        VALUES ($1, $2, $3) RETURNING id
    `, user.FirstName, user.LastName, user.Age).Scan(&id)

	return id, err
}

func Get(db *sql.DB, id string) (*User, error) {
	rows, err := db.Query(`
		SELECT u.id, u.first_name, u.last_name, u.age, c.id, c.name, c.model, c.year FROM data.users u
		LEFT JOIN data.users_cars uc ON u.id = uc.user_id
	  	LEFT JOIN data.cars c ON uc.car_id = c.id
		WHERE u.id = $1`,
		id,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var u *User
	for rows.Next() {
		var (
			userID              int
			firstName, lastName string
			age                 int
			carID               *int
			carName, carModel   *string
			carYear             *int
		)

		if err := rows.Scan(&userID, &firstName, &lastName, &age, &carID, &carName, &carModel, &carYear); err != nil {
			return nil, err
		}

		carExists := carID != nil && carName != nil && carModel != nil && carYear != nil

		car := func() []users_cars.CarWithoutUser {
			if carExists {
				return []users_cars.CarWithoutUser{{
					ID:    *carID,
					Name:  *carName,
					Model: *carModel,
					Year:  *carYear,
				}}
			}
			return nil
		}()

		if u == nil {
			u = &User{
				ID:        userID,
				FirstName: firstName,
				LastName:  lastName,
				Age:       age,
				Cars:      car,
			}
		} else {
			if carExists {
				u.Cars = append(u.Cars, car[0])
			}
		}
	}

	return u, nil
}

func GetAll(db *sql.DB) ([]User, error) {
	rows, err := db.Query(`
		SELECT u.id, u.first_name, u.last_name, u.age, c.id, c.name, c.model, c.year FROM data.users u
		LEFT JOIN data.users_cars uc ON u.id = uc.user_id
	  	LEFT JOIN data.cars c ON uc.car_id = c.id
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := map[int]User{}
	for rows.Next() {
		var (
			userID              int
			firstName, lastName string
			age                 int
			carID               *int
			carName, carModel   *string
			carYear             *int
		)

		if err := rows.Scan(&userID, &firstName, &lastName, &age, &carID, &carName, &carModel, &carYear); err != nil {
			return nil, err
		}

		carExists := carID != nil && carName != nil && carModel != nil && carYear != nil

		car := func() []users_cars.CarWithoutUser {
			if carExists {
				return []users_cars.CarWithoutUser{{
					ID:    *carID,
					Name:  *carName,
					Model: *carModel,
					Year:  *carYear,
				}}
			}
			return nil
		}()

		if user, ok := users[userID]; ok {
			if carExists {
				user.Cars = append(user.Cars, car[0])
				users[userID] = user
			}
		} else {
			users[userID] = User{
				ID:        userID,
				FirstName: firstName,
				LastName:  lastName,
				Age:       age,
				Cars:      car,
			}
		}
	}

	return slices.Collect(maps.Values(users)), nil
}
