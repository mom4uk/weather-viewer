package testutils

import (
	"database/sql"
	"log"
	"os"
)

type TestDB struct {
	DB *sql.DB
}

func NewTestDB() *TestDB {
	dsl := os.Getenv("DATABASE_URL")

	database, err := sql.Open("postgres", dsl)
	if err != nil {
		log.Fatal(err)
	}

	if err := database.Ping(); err != nil {
		log.Fatal(err)
	}

	return &TestDB{
		DB: database,
	}
}
