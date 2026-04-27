package db

import (
	"database/sql"
	"os"
)

func InitDB() *sql.DB {
	dsl := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", dsl)
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}

	return db
}
