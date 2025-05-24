package car

type Car struct {
	ID    int
	Name  string
	Model string
	Year  int
}

type CreateCarRequest struct {
	Name  string `json:"name"`
	Model string `json:"model"`
	Year  int    `json:"year"`
}

type Dto struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Model string `json:"model"`
	Year  int    `json:"year"`
}

func AsDto(car Car) Dto {
	return Dto{
		ID:    car.ID,
		Name:  car.Name,
		Model: car.Model,
		Year:  car.Year,
	}
}
