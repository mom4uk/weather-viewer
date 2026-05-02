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

	//userService := services.NewUserService(userRepository)
	locationService := services.NewLocationService(locationRepository)
	sessionService := services.NewSessionService(sessionRepository)

	//userController := controllers.NewUserController()
	locationController := controllers.NewLocationController(locationService)

	//controllers.RegisterUserRoutes(srv.GetMux(), userController)
	controllers.RegisterLocationRoutes(srv.GetMux(), locationController, sessionService)
	return &TestApp{
		DB:     db,
		Server: srv,
	}
}
