package dto

import (
	"strings"
	"weather-viewer/internal/domain"
)

type LocationResponse struct {
	ID        int
	Name      string
	UserID    int
	Latitude  float64
	Longitude float64
}

func ValidateId(id string) error {
	id = strings.TrimSpace(id)
	if id == "" {
		return domain.ErrInvalidId
	}
	return nil
}
