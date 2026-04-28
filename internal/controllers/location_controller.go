package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"weather-viewer/internal/apierrors"
	"weather-viewer/internal/domain"
	"weather-viewer/internal/dto"
	"weather-viewer/internal/services"
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
	idStr := r.PathValue("id")
	if err := dto.ValidateId(idStr); err != nil {
		apierrors.HandleError(w, err)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		apierrors.HandleError(w, domain.ErrInvalidId)
		return
	}

	result, err := c.locationService.GetLocation(id)
	if err != nil {
		apierrors.HandleError(w, err)
		return
	}

	response := dto.LocationResponse{
		ID:        result.ID,
		Name:      result.Name,
		UserID:    result.UserID,
		Latitude:  result.Latitude,
		Longitude: result.Longitude,
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		apierrors.WriteError(w, "Ошибка при формировании json", http.StatusInternalServerError)
	}
}
