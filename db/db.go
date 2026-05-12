package db

import (
	"context"
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
)

func InitPostgres() *sql.DB {
	dsl := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", dsl)
	log.Println("DATABASE_URL =", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}

	for i := 0; i < 10; i++ {
		err := db.Ping()
		if err == nil {
			break
		}
		time.Sleep(2 * time.Second)
	}

	return db
}

func InitRedis() *redis.Client {
	addr := os.Getenv("REDIS_URL")
	log.Println("REDIS_URL =", os.Getenv("REDIS_URL"))
	opt, err := redis.ParseURL(addr)
	if err != nil {
		panic(err)
	}

	client := redis.NewClient(opt)
	if err := client.Ping(context.Background()).Err(); err != nil {
		panic(err)
	}
	return client
}
