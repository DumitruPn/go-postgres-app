package user

import "go-postgres-app/internal/car"

type User struct {
	ID        int
	FirstName string
	LastName  string
	Age       int
	Car       []car.Car
}

type CreateUserRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Age       int    `json:"age"`
}

type Dto struct {
	ID        int       `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Age       int       `json:"age"`
	Car       []car.Dto `json:"car"`
}

func AsDto(user User) Dto {
	return Dto{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Age:       user.Age,
		Car:       []car.Dto{},
	}
}
