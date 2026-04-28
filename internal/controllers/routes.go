package controllers

import (
	"net/http"
	"weather-viewer/internal/middlewares"
)

func RegisterUserRoutes(mux *http.ServeMux, c *UserController) {
	mux.Handle("POST /auth/register", middlewares.JSON(http.HandlerFunc(c.RegisterUser)))
	mux.Handle("POST /auth/login", middlewares.JSON(http.HandlerFunc(c.LoginUser)))
	mux.Handle("POST /auth/logout", middlewares.JSON(http.HandlerFunc(c.LogoutUser)))
}

func RegisterLocationRoutes(mux *http.ServeMux, c *LocationsController) {
	mux.Handle("GET /searchLocation", middlewares.JSON(http.HandlerFunc(c.GetLocation)))
	mux.Handle("POST /addLocation", middlewares.JSON(http.HandlerFunc(c.AddLocation)))
	mux.Handle("GET /getLocations", middlewares.JSON(http.HandlerFunc(c.GetLocations)))
	mux.Handle("DELETE /removeLocation", middlewares.JSON(http.HandlerFunc(c.RemoveLocation)))
}
