package main

import (
	"log"
	"weather-viewer/db"
	"weather-viewer/server"
)

func main() {
	db, err := db.InitDB()
	if err != nil {
		panic(err)
	}

	srv := server.NewServer()
	if err := srv.Start(); err != nil {
		log.Fatal(err)
	}
}
