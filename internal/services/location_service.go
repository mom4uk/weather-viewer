package services

import (
	"weather-viewer/internal/domain"
	"weather-viewer/internal/ports"
)

type LocationService struct {
	locationRepository ports.LocationRepository
}

func NewLocationService(locationRepository ports.LocationRepository) *LocationService {
	return &LocationService{
		locationRepository: locationRepository,
	}
}

func (s *LocationService) GetLocation(id int) (domain.Location, error) {
	return s.locationRepository.GetLocation(id)
}

func (s *LocationService) AddLocation(location domain.Location) (domain.Location, error) {
	return s.locationRepository.AddLocation(location)
}
