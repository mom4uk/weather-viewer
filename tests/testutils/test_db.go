package testutils

import (
	"context"
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
)

type TestDB struct {
	Postgres *sql.DB
	Redis    *redis.Client
}

func NewTestDB() *TestDB {
	return &TestDB{
		Postgres: NewTestPostgres(),
		Redis:    NewTestRedis(),
	}
}

func NewTestPostgres() *sql.DB {
	dsl := os.Getenv("DATABASE_URL")
	if dsl == "" {
		panic("DATABASE_URL environment variable not set")
	}
	database, err := sql.Open("postgres", dsl)
	if err != nil {
		log.Fatal(err)
	}

	if err := database.Ping(); err != nil {
		log.Fatal(err)
	}

	return database
}

func NewTestRedis() *redis.Client {
	addr := os.Getenv("REDIS_ADDR")
	if addr == "" {
		panic("REDIS_ADDR environment variable not set")
	}
	client := redis.NewClient(&redis.Options{
		Addr: addr,
	})
	if err := client.Ping(context.Background()).Err(); err != nil {
		panic(err)
	}
	return client
}
