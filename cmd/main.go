package main

import (
	"log"
	"weather-viewer/db"
	"weather-viewer/internal/controllers"
	"weather-viewer/internal/middlewares"
	"weather-viewer/internal/repositories"
	"weather-viewer/internal/services"
	"weather-viewer/server"
)

func main() {
	database := db.InitDB()

	srv := server.NewServer()

	userRepository := repositories.NewUserRepository(database)
	locationRepository := repositories.NewLocationRepository(database)
	sessionRepository := repositories.NewSessionRepository(database)

	userService := services.NewUserService(userRepository)
	sessionService := services.NewSessionService(sessionRepository)
	locationService := services.NewLocationService(locationRepository)
	authService := services.NewAuthService(sessionService, userService)

	middlewares.Auth(sessionService)

	userController := controllers.NewAuthController(userService, sessionService, authService)
	locationController := controllers.NewLocationController(locationService)

	controllers.RegisterAuthRoutes(srv.GetMux(), userController)
	controllers.RegisterLocationRoutes(srv.GetMux(), locationController, sessionService)

	if err := srv.Start(); err != nil {
		log.Fatal(err)
	}
}
