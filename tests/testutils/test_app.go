package testutils

import (
	"database/sql"
	"net/http"
	"os"
	"weather-viewer/internal/clients"
	"weather-viewer/internal/controllers"
	"weather-viewer/internal/interfaces"
	"weather-viewer/internal/render"
	"weather-viewer/internal/repositories"
	"weather-viewer/internal/services"
	"weather-viewer/server"

	"github.com/redis/go-redis/v9"
)

type TestApp struct {
	Postgres      *sql.DB
	Server        *server.Server
	Redis         *redis.Client
	WeatherClient interfaces.Weather
}

func NewTestApp(db *TestDB) *TestApp {
	srv := server.NewServer()

	locationRepository := repositories.NewLocationRepository(db.Postgres)
	sessionRepository := repositories.NewSessionRepository(db.Redis)
	userRepository := repositories.NewUserRepository(db.Postgres)

	apiKey := os.Getenv("WEATHER_API_KEY")
	weatherClient := clients.NewWeatherClient("https://api.openweathermap.org", apiKey, http.DefaultClient)

	userService := services.NewUserService(userRepository)
	locationService := services.NewLocationService(locationRepository, weatherClient)
	sessionService := services.NewSessionService(sessionRepository)
	authService := services.NewAuthService(sessionService, userService)
	renderer, err := render.NewTemplateRenderer()
	if err != nil {
		panic(err)
	}

	userController := controllers.NewAuthController(userService, sessionService, authService, renderer)
	locationController := controllers.NewLocationController(locationService)

	controllers.RegisterAuthRoutes(srv.GetMux(), userController)
	controllers.RegisterLocationRoutes(srv.GetMux(), locationController, sessionService)
	return &TestApp{
		Postgres: db.Postgres,
		Redis:    db.Redis,
		Server:   srv,
	}
}

func NewTestAppForWeather(db *TestDB, weatherClient interfaces.Weather) *TestApp {
	srv := server.NewServer()

	locationRepository := repositories.NewLocationRepository(db.Postgres)
	sessionRepository := repositories.NewSessionRepository(db.Redis)
	userRepository := repositories.NewUserRepository(db.Postgres)

	userService := services.NewUserService(userRepository)
	locationService := services.NewLocationService(locationRepository, weatherClient)
	sessionService := services.NewSessionService(sessionRepository)
	authService := services.NewAuthService(sessionService, userService)
	renderer, err := render.NewTemplateRenderer()
	if err != nil {
		panic(err)
	}

	userController := controllers.NewAuthController(userService, sessionService, authService, renderer)
	locationController := controllers.NewLocationController(locationService)

	controllers.RegisterAuthRoutes(srv.GetMux(), userController)
	controllers.RegisterLocationRoutes(srv.GetMux(), locationController, sessionService)
	return &TestApp{
		Postgres: db.Postgres,
		Redis:    db.Redis,
		Server:   srv,
	}
}
