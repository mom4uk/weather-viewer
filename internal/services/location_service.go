package services

import (
	"weather-viewer/internal/domain"
	"weather-viewer/internal/interfaces"
)

type LocationService struct {
	LocationRepository interfaces.LocationRepository
	WeatherClient      interfaces.Weather
}

func NewLocationService(locationRepository interfaces.LocationRepository, WeatherClient interfaces.Weather) *LocationService {
	return &LocationService{
		LocationRepository: locationRepository,
		WeatherClient:      WeatherClient,
	}
}

func (s *LocationService) GetLocation(id int) (domain.Location, error) {
	location, err := s.LocationRepository.GetLocation(id)
	if err != nil {
		return domain.Location{}, err
	}
	weather, err := s.WeatherClient.GetWeather(location)
	location.Weather = weather
	return location, err
}

func (s *LocationService) AddLocation(location domain.Location) (domain.Location, error) {
	return s.LocationRepository.AddLocation(location)
}

func (s *LocationService) GetLocations(userID int) ([]domain.Location, error) {
	locations, err := s.LocationRepository.GetLocations(userID)
	if err != nil {
		return []domain.Location{}, err
	}
	for i := range locations {
		weather, err := s.WeatherClient.GetWeather(locations[i])
		if err != nil {
			return []domain.Location{}, err
		}
		locations[i].Weather = weather
	}
	return locations, err
}

func (s *LocationService) RemoveLocation(id int) error {
	return s.LocationRepository.RemoveLocation(id)
}
