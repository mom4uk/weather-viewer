package controllers

import (
	"encoding/json"
	"net/http"
	"weather-viewer/internal/apierrors"
	"weather-viewer/internal/httputil"
)

type LocationController struct {
	locationService *services.LocationService
}

func NewLocationController(s *services.LocationService) *LocationController {
	return &LocationController{
		locationService: s,
	}
}

func (c *LocationController) GetLocation(w http.ResponseWriter, r *http.Request) {
	id, err := httputil.GetIdFromUrl(r.URL.Path)
	if err != nil {
		apierrors.HandleError(w, err)
		return
	}

	result, err := services.GetLocation(c.locationService, id)
	if err != nil {
		apierrors.HandleError(w, err)
		return
	}

	if err := json.NewEncoder(w).Encode(result); err != nil {
		apierrors.WriteError(w, "Ошибка при формировании json", http.StatusInternalServerError)
	}
}
