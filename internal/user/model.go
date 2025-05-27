package user

import "go-postgres-app/internal/users_cars"

type User struct {
	ID        int
	FirstName string
	LastName  string
	Age       int
	Cars      []users_cars.CarWithoutUser
}

type CreateUserRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Age       int    `json:"age"`
}

type Dto struct {
	ID        int                            `json:"id"`
	FirstName string                         `json:"firstName"`
	LastName  string                         `json:"lastName"`
	Age       int                            `json:"age"`
	Cars      []users_cars.DtoCarWithoutUser `json:"cars"`
}

func AsDtoWithCars(user User) Dto {
	return Dto{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Age:       user.Age,
		Cars:      users_cars.AsDtosWithoutUsers(user.Cars),
	}
}

func AsDtosWithCars(users []User) []Dto {
	dtos := []Dto{}
	for _, user := range users {
		dtos = append(dtos, AsDtoWithCars(user))
	}
	return dtos
}
