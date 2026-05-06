package testutils

import (
	"weather-viewer/internal/controllers"
	"weather-viewer/internal/repositories"
	"weather-viewer/internal/services"
	"weather-viewer/server"
)

type TestApp struct {
	DB     *TestDB
	Server *server.Server
}

func NewTestApp(db *TestDB) *TestApp {
	srv := server.NewServer()

	locationRepository := repositories.NewLocationRepository(db.DB)
	sessionRepository := repositories.NewSessionRepository(db.DB)
	userRepository := repositories.NewUserRepository(db.DB)

	userService := services.NewUserService(userRepository)
	locationService := services.NewLocationService(locationRepository)
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
