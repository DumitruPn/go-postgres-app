package user

import "go-postgres-app/internal/car"

type User struct {
	ID        int
	FirstName string
	LastName  string
	Age       int
	Cars      []car.Car
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
	Cars      []car.Dto `json:"cars"`
}

func AsDto(user User) Dto {
	return Dto{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Age:       user.Age,
		Cars:      car.AsDtos(user.Cars),
	}
}

func AsDtos(users []User) []Dto {
	dtos := []Dto{}
	for _, user := range users {
		dtos = append(dtos, AsDto(user))
	}
	return dtos
}
