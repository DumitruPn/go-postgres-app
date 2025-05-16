package models

type Car struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Model string `json:"model"`
    Year  int    `json:"year"`
}
