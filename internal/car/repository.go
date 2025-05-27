package car

import (
	"database/sql"
	"go-postgres-app/internal/users_cars"
	"maps"
	"slices"
)

func Insert(db *sql.DB, car CreateCarRequest) (int, error) {
	var id int
	err := db.QueryRow(`
		INSERT INTO data.cars (name, model, year)
		VALUES ($1, $2, $3) RETURNING id
	`, car.Name, car.Model, car.Year).Scan(&id)

	return id, err
}

func Get(db *sql.DB, id string) (*Car, error) {
	rows, err := db.Query(`
		SELECT c.id, c.name, c.model, c.year, u.id, u.first_name, u.last_name, u.age  FROM data.cars c
		LEFT JOIN data.users_cars uc ON c.id = uc.car_id
		LEFT JOIN data.users u ON uc.user_id = u.id
		WHERE c.id = $1`,
		id,
	)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var c *Car
	for rows.Next() {
		var (
			carID                       int
			name, model                 string
			year                        int
			userID                      *int
			userFirstName, userLastName *string
			userAge                     *int
		)

		if err := rows.Scan(&carID, &name, &model, &year, &userID, &userFirstName, &userLastName, &userAge); err != nil {
			return nil, err
		}

		userExists := userID != nil && userFirstName != nil && userLastName != nil && userAge != nil

		user := func() []users_cars.UserWithoutCar {
			if userExists {
				return []users_cars.UserWithoutCar{{
					ID:        *userID,
					FirstName: *userFirstName,
					LastName:  *userLastName,
					Age:       *userAge,
				}}
			}
			return nil
		}()

		if c == nil {
			c = &Car{
				ID:    carID,
				Name:  name,
				Model: model,
				Year:  year,
				Users: user,
			}
		} else {
			if userExists {
				c.Users = append(c.Users, user[0])
			}
		}
	}

	return c, nil
}

func GetAll(db *sql.DB) ([]Car, error) {
	rows, err := db.Query(`
		SELECT c.id, c.name, c.model, c.year, u.id, u.first_name, u.last_name, u.age FROM data.cars c
		LEFT JOIN data.users_cars uc ON c.id = uc.car_id
	  	LEFT JOIN data.users u ON uc.user_id = u.id
	`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	cars := map[int]Car{}
	for rows.Next() {
		var (
			carID                       int
			name, model                 string
			year                        int
			userID                      *int
			userFirstName, userLastName *string
			userAge                     *int
		)

		if err := rows.Scan(&carID, &name, &model, &year, &userID, &userFirstName, &userLastName, &userAge); err != nil {
			return nil, err
		}

		userExists := userID != nil && userFirstName != nil && userLastName != nil && userAge != nil

		user := func() []users_cars.UserWithoutCar {
			if userExists {
				return []users_cars.UserWithoutCar{{
					ID:        *userID,
					FirstName: *userFirstName,
					LastName:  *userLastName,
					Age:       *userAge,
				}}
			}
			return nil
		}()

		if car, ok := cars[carID]; ok {
			if userExists {
				car.Users = append(car.Users, user[0])
				cars[carID] = car
			}
		} else {
			cars[carID] = Car{
				ID:    carID,
				Name:  name,
				Model: model,
				Year:  year,
				Users: user,
			}
		}
	}
	return slices.Collect(maps.Values(cars)), nil
}
