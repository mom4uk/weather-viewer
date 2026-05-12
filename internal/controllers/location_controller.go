package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"weather-viewer/internal/apierrors"
	"weather-viewer/internal/contextkeys"
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
	// тут надо user_id достать и проверять, что я не могу с другого юзера почекать локации?
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
		Weather:   result.Weather,
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		apierrors.WriteError(w, "Ошибка при формировании json", http.StatusInternalServerError)
	}
}

func (c *LocationController) AddLocation(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(contextkeys.UserID).(int)

	if !ok {
		apierrors.WriteError(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	lat, err := strconv.ParseFloat(r.FormValue("latitude"), 64)
	if err != nil {
		apierrors.HandleError(w, domain.ErrInvalidLatitude)
		return
	}

	lon, err := strconv.ParseFloat(r.FormValue("longitude"), 64)
	if err != nil {
		apierrors.HandleError(w, domain.ErrInvalidLongitude)
		return
	}

	if err := dto.ValidateName(r.FormValue("name")); err != nil {
		apierrors.HandleError(w, err)
		return
	}

	result, err := c.locationService.AddLocation(domain.Location{
		Name:      r.FormValue("name"),
		UserID:    userID,
		Latitude:  lat,
		Longitude: lon,
	})
	if err != nil {
		apierrors.HandleError(w, err)
		return
	}

	if strings.Contains(r.Header.Get("Accept"), "text/html") {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	w.WriteHeader(http.StatusCreated)

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

func (c *LocationController) GetLocations(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(contextkeys.UserID).(int)
	if !ok {
		apierrors.WriteError(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	result, err := c.locationService.GetLocations(userID)
	if err != nil {
		apierrors.HandleError(w, err)
		return
	}

	var response []dto.LocationResponse

	for _, location := range result {
		response = append(response, dto.LocationResponse{
			ID:        location.ID,
			Name:      location.Name,
			UserID:    location.UserID,
			Latitude:  location.Latitude,
			Longitude: location.Longitude,
			Weather:   location.Weather,
		})
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		apierrors.WriteError(w, "Ошибка при формировании json", http.StatusInternalServerError)
	}
}

func (c *LocationController) RemoveLocation(w http.ResponseWriter, r *http.Request) {
	locationIDStr := r.PathValue("id")
	locationID, err := strconv.Atoi(locationIDStr)
	if err != nil {
		apierrors.HandleError(w, domain.ErrInvalidId)
		return
	}
	if err = c.locationService.RemoveLocation(locationID); err != nil {
		apierrors.HandleError(w, err)
		return
	}

	if strings.Contains(r.Header.Get("Accept"), "text/html") {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	w.WriteHeader(http.StatusNoContent)

	if err := json.NewEncoder(w).Encode(nil); err != nil {
		apierrors.WriteError(w, "Ошибка при формировании json", http.StatusInternalServerError)
	}
}
