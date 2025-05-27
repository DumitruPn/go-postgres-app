package car

import "go-postgres-app/internal/users_cars"

type Car struct {
	ID    int
	Name  string
	Model string
	Year  int
	Users []users_cars.UserWithoutCar
}

type CreateCarRequest struct {
	Name  string `json:"name"`
	Model string `json:"model"`
	Year  int    `json:"year"`
}

type Dto struct {
	ID    int                            `json:"id"`
	Name  string                         `json:"name"`
	Model string                         `json:"model"`
	Year  int                            `json:"year"`
	Users []users_cars.DtoUserWithoutCar `json:"users"`
}

func AsDtoWithUsers(car Car) Dto {
	return Dto{
		ID:    car.ID,
		Name:  car.Name,
		Model: car.Model,
		Year:  car.Year,
		Users: users_cars.AsDtosWithoutCars(car.Users),
	}
}

func AsDtosWithUsers(cars []Car) []Dto {
	dtos := []Dto{}
	for _, car := range cars {
		dtos = append(dtos, AsDtoWithUsers(car))
	}
	return dtos
}
