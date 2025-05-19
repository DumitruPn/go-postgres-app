package user

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Age       int    `json:"age"`
	CarID     int    `json:"carId"`
}
