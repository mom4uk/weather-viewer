package main

import (
	"log"
	"weather-viewer/db"
	"weather-viewer/internal/controllers"
	"weather-viewer/server"
)

func main() {
	db, err := db.InitDB()
	if err != nil {
		panic(err)
	}

	srv := server.NewServer()

	userController := controllers.NewUserController()
	locationController := controllers.NewLocationController()

	userService := services.NewUserService(userController)
	locationService := services.NewLocationService(locationController)

	userRepository := repositories.NewUserRepository(db)
	locationRepository := repositories.NewLocationRepository(db)

	controllers.RegisterUserRoutes(srv.GetMux(), userController)
	controllers.RegisterLocationRoutes(srv.GetMux(), locationController)

	if err := srv.Start(); err != nil {
		log.Fatal(err)
	}
}
