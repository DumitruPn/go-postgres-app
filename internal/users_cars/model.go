package users_cars

// User
type UserWithoutCar struct {
	ID        int
	FirstName string
	LastName  string
	Age       int
}

type DtoUserWithoutCar struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Age       int    `json:"age"`
}

func AsDtoWithoutCars(user UserWithoutCar) DtoUserWithoutCar {
	return DtoUserWithoutCar{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Age:       user.Age,
	}
}

func AsDtosWithoutCars(users []UserWithoutCar) []DtoUserWithoutCar {
	dtos := []DtoUserWithoutCar{}
	for _, user := range users {
		dtos = append(dtos, AsDtoWithoutCars(user))
	}
	return dtos
}

// Car
type CarWithoutUser struct {
	ID    int
	Name  string
	Model string
	Year  int
}

type DtoCarWithoutUser struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Model string `json:"model"`
	Year  int    `json:"year"`
}

func AsDtoWithoutUsers(car CarWithoutUser) DtoCarWithoutUser {
	return DtoCarWithoutUser{
		ID:    car.ID,
		Name:  car.Name,
		Model: car.Model,
		Year:  car.Year,
	}
}

func AsDtosWithoutUsers(cars []CarWithoutUser) []DtoCarWithoutUser {
	dtos := []DtoCarWithoutUser{}
	for _, car := range cars {
		dtos = append(dtos, AsDtoWithoutUsers(car))
	}
	return dtos
}
