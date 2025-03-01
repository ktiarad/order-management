package repository

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func ConnectDB() *sql.DB {
	godotenv.Load()

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DBHost"),
		os.Getenv("DBPort"),
		os.Getenv("DBUser"),
		os.Getenv("DBPass"),
		os.Getenv("DBName"))

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	return db
}
