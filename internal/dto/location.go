package dto

import (
	"regexp"
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

type AddLocationRequest struct {
	Name      string
	ID        int
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

func ValidateName(name string) error {
	nameTrimed := strings.TrimSpace(name)
	var nameRegex = regexp.MustCompile(`^[\p{L}\s-]+$`)

	if nameTrimed == "" {
		return domain.ErrInvalidName
	}

	if !nameRegex.MatchString(nameTrimed) {
		return domain.ErrInvalidName
	}
	return nil
}
