package main

import (
	"log"
	"net/http"
	"os"
	"weather-viewer/db"
	"weather-viewer/internal/clients"
	"weather-viewer/internal/controllers"
	"weather-viewer/internal/middlewares"
	"weather-viewer/internal/render"
	"weather-viewer/internal/repositories"
	"weather-viewer/internal/services"
	"weather-viewer/server"
)

func main() {
	postgres := db.InitPostgres()
	redis := db.InitRedis()

	srv := server.NewServer()
	srv.InitStatic()

	userRepository := repositories.NewUserRepository(postgres)
	locationRepository := repositories.NewLocationRepository(postgres)
	sessionRepository := repositories.NewSessionRepository(redis)

	apiKey := os.Getenv("WEATHER_API_KEY")
	weatherClient := clients.NewWeatherClient("https://api.openweathermap.org", apiKey, http.DefaultClient)

	userService := services.NewUserService(userRepository)
	sessionService := services.NewSessionService(sessionRepository)
	locationService := services.NewLocationService(locationRepository, weatherClient)
	authService := services.NewAuthService(sessionService, userService)

	middlewares.Auth(sessionService)

	renderer, err := render.NewTemplateRenderer()
	if err != nil {
		log.Fatal(err)
	}

	pageController := controllers.NewPageController(
		renderer,
		locationService,
		sessionService,
	)
	userController := controllers.NewAuthController(userService, sessionService, authService)
	locationController := controllers.NewLocationController(locationService)

	controllers.RegisterPageRoutes(srv.GetMux(), pageController)
	controllers.RegisterAuthRoutes(srv.GetMux(), userController)
	controllers.RegisterLocationRoutes(srv.GetMux(), locationController, sessionService)

	if err := srv.Start(); err != nil {
		log.Fatal(err)
	}
}
