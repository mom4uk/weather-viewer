package controllers

import "net/http"

func RegisterUserRoutes(mux *http.ServeMux, c *UserController) {
	mux.Handle("POST /auth/register", http.HandlerFunc(c.RegisterUser))
	mux.Handle("POST /auth/login", http.HandlerFunc(c.LoginUser))
	mux.Handle("POST /auth/logout", http.HandlerFunc(c.LogoutUser))
}

func RegisterLocationRoutes(mux *http.ServeMux, c *LocationsController) {
	mux.Handle("GET /searchLocation", http.HandlerFunc(c.GetLocation))
	mux.Handle("POST /addLocation", http.HandlerFunc(c.AddLocation))
	mux.Handle("GET /getLocations", http.HandlerFunc(c.GetLocations))
	mux.Handle("DELETE /removeLocation", http.HandlerFunc(c.RemoveLocation))
}
