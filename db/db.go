package db

import (
	"context"
	"database/sql"
	"os"

	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
)

func InitPostgres() *sql.DB {
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

func InitRedis() *redis.Client {
	addr := os.Getenv("REDIS_ADDR")
	client := redis.NewClient(&redis.Options{
		Addr: addr,
	})
	if err := client.Ping(context.Background()).Err(); err != nil {
		panic(err)
	}
	return client
}
