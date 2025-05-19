package main

import (
	"fmt"
	"log"
	"net/http"

	"go-postgres-app/internal/db"
	"go-postgres-app/internal/routes"
)

func main() {
	database := db.Connect()

	router := routes.Setup(database)

	fmt.Println("🚀 Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
