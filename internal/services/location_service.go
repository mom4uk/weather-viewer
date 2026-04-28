package services

import (
	"weather-viewer/internal/domain"
	"weather-viewer/internal/repositories"
)

type LocationService struct {
	locationRepository *repositories.LocationRepository
}

func NewLocationService(locationRepository *repositories.LocationRepository) *LocationService {
	return &LocationService{
		locationRepository: locationRepository,
	}
}

func (s *LocationService) GetLocation(id int) (domain.Location, error) {
	return s.locationRepository.GetLocation(id)
}
