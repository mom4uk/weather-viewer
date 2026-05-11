package controllers

import (
	"net/http"
	"weather-viewer/internal/middlewares"
	"weather-viewer/internal/services"
)

func RegisterAuthRoutes(mux *http.ServeMux, c *AuthController) {
	mux.Handle(
		"POST /auth/register",
		middlewares.Chain(
			middlewares.JSON(),
		)(http.HandlerFunc(c.SignUp)),
	)
	mux.Handle(
		"POST /auth/login",
		middlewares.Chain(
			middlewares.JSON(),
		)(http.HandlerFunc(c.SignIn)))
	mux.Handle(
		"POST /auth/logout",
		middlewares.Chain(
			middlewares.JSON(),
		)(http.HandlerFunc(c.SignOut)))
}

func RegisterLocationRoutes(mux *http.ServeMux, c *LocationController, s *services.SessionService) {
	mux.Handle(
		"GET /searchLocation/{id}",
		middlewares.Chain(
			middlewares.Auth(s),
			middlewares.JSON(),
		)(http.HandlerFunc(c.GetLocation)),
	)
	mux.Handle(
		"POST /addLocation",
		middlewares.Chain(
			middlewares.Auth(s),
			middlewares.JSON(),
		)(http.HandlerFunc(c.AddLocation)),
	)
	mux.Handle(
		"GET /getLocations",
		middlewares.Chain(
			middlewares.Auth(s),
			middlewares.JSON(),
		)(http.HandlerFunc(c.GetLocations)),
	)
	mux.Handle("DELETE /removeLocation/{id}",
		middlewares.Chain(
			middlewares.Auth(s),
			middlewares.JSON(),
		)(http.HandlerFunc(c.RemoveLocation)),
	)
}
