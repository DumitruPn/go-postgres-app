package db

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

func Connect() *sql.DB {
	_, b, _, _ := runtime.Caller(0)
	basePath := filepath.Dir(b)
	envPath := filepath.Join(basePath, "../../.env")

	err := godotenv.Load(envPath)
	if err != nil {
		log.Fatalf("Error loading .env file from %s", envPath)
	}

	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
	)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`
        CREATE SCHEMA IF NOT EXISTS data;

        CREATE TABLE IF NOT EXISTS data.cars (
            id SERIAL PRIMARY KEY,
            name TEXT NOT NULL,
            model TEXT NOT NULL,
            year INT NOT NULL
        );

        CREATE TABLE IF NOT EXISTS data.users (
            id SERIAL PRIMARY KEY,
            first_name TEXT NOT NULL,
            last_name TEXT NOT NULL,
            age INT NOT NULL
        );

		CREATE TABLE IF NOT EXISTS data.users_cars (
		    user_id INT REFERENCES data.users(id),
		    car_id INT REFERENCES data.cars(id),
		    
		    PRIMARY KEY (user_id, car_id)
		);
		
		CREATE TABLE IF NOT EXISTS data.notifications (
		    id SERIAL PRIMARY KEY,
		    value TEXT NOT NULL
		)
    `)
	if err != nil {
		log.Fatal("Failed to ensure cars table:", err)
	}

	return db
}
