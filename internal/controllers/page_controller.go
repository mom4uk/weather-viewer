package controllers

import (
	"net/http"
	"weather-viewer/internal/domain"
	"weather-viewer/internal/render"
	"weather-viewer/internal/services"
)

type PageController struct {
	renderer        *render.TemplateRenderer
	locationService *services.LocationService
	sessionService  *services.SessionService
}

type HomePageData struct {
	Locations []domain.Location
}

func NewPageController(
	renderer *render.TemplateRenderer,
	locationService *services.LocationService,
	sessionService *services.SessionService,
) *PageController {
	return &PageController{
		renderer:        renderer,
		locationService: locationService,
		sessionService:  sessionService,
	}
}

func (c *PageController) Home(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_token")
	if err != nil {
		http.Redirect(w, r, "/sign-in", http.StatusSeeOther)
		return
	}

	userID, err := c.sessionService.GetUserID(r.Context(), cookie.Value)
	if err != nil {
		http.Redirect(w, r, "/sign-in", http.StatusSeeOther)
		return
	}

	locations, err := c.locationService.GetLocations(userID)
	if err != nil {
		c.renderer.Render(w, "error.html", nil)
		return
	}

	c.renderer.Render(w, "index.html", HomePageData{Locations: locations})
}

func (c *PageController) SignIn(w http.ResponseWriter, _ *http.Request) {
	c.renderer.Render(w, "sign-in.html", nil)
}

func (c *PageController) SignUp(w http.ResponseWriter, _ *http.Request) {
	c.renderer.Render(w, "sign-up.html", nil)
}

func (c *PageController) Error(w http.ResponseWriter, _ *http.Request) {
	c.renderer.Render(w, "error.html", nil)
}
