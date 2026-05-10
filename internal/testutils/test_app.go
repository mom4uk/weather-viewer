package testutils

import (
	"net/http"
	"os"
	"weather-viewer/internal/clients"
	"weather-viewer/internal/controllers"
	"weather-viewer/internal/interfaces"
	"weather-viewer/internal/repositories"
	"weather-viewer/internal/services"
	"weather-viewer/server"
)

type TestApp struct {
	DB            *TestDB
	Server        *server.Server
	WeatherClient interfaces.Weather
}

func NewTestApp(db *TestDB) *TestApp {
	srv := server.NewServer()

	locationRepository := repositories.NewLocationRepository(db.DB)
	sessionRepository := repositories.NewSessionRepository(db.DB)
	userRepository := repositories.NewUserRepository(db.DB)

	apiKey := os.Getenv("WEATHER_API_KEY")
	weatherClient := clients.NewWeatherClient("https://api.openweathermap.org", apiKey, http.DefaultClient)

	userService := services.NewUserService(userRepository)
	locationService := services.NewLocationService(locationRepository, weatherClient)
	sessionService := services.NewSessionService(sessionRepository)
	authService := services.NewAuthService(sessionService, userService)

	userController := controllers.NewAuthController(userService, sessionService, authService)
	locationController := controllers.NewLocationController(locationService)

	controllers.RegisterAuthRoutes(srv.GetMux(), userController)
	controllers.RegisterLocationRoutes(srv.GetMux(), locationController, sessionService)
	return &TestApp{
		DB:     db,
		Server: srv,
	}
}

func NewTestAppForWeather(db *TestDB, weatherClient interfaces.Weather) *TestApp {
	srv := server.NewServer()

	locationRepository := repositories.NewLocationRepository(db.DB)
	sessionRepository := repositories.NewSessionRepository(db.DB)
	userRepository := repositories.NewUserRepository(db.DB)

	userService := services.NewUserService(userRepository)
	locationService := services.NewLocationService(locationRepository, weatherClient)
	sessionService := services.NewSessionService(sessionRepository)
	authService := services.NewAuthService(sessionService, userService)

	userController := controllers.NewAuthController(userService, sessionService, authService)
	locationController := controllers.NewLocationController(locationService)

	controllers.RegisterAuthRoutes(srv.GetMux(), userController)
	controllers.RegisterLocationRoutes(srv.GetMux(), locationController, sessionService)
	return &TestApp{
		DB:     db,
		Server: srv,
	}
}
