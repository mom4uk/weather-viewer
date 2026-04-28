package main

import (
	"log"
	"weather-viewer/db"
	"weather-viewer/internal/controllers"
	"weather-viewer/internal/repositories"
	"weather-viewer/internal/services"
	"weather-viewer/server"
)

func main() {
	database := db.InitDB()

	srv := server.NewServer()

	//userRepository := repositories.NewUserRepository(db)
	locationRepository := repositories.NewLocationRepository(database)

	//userService := services.NewUserService(userRepository)
	locationService := services.NewLocationService(locationRepository)

	//userController := controllers.NewUserController()
	locationController := controllers.NewLocationController(locationService)

	//controllers.RegisterUserRoutes(srv.GetMux(), userController)
	controllers.RegisterLocationRoutes(srv.GetMux(), locationController)

	if err := srv.Start(); err != nil {
		log.Fatal(err)
	}
}
