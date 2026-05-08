package services

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"weather-viewer/internal/domain"
	"weather-viewer/internal/interfaces"
)

type LocationService struct {
	locationRepository interfaces.LocationRepository
}

func NewLocationService(locationRepository interfaces.LocationRepository) *LocationService {
	return &LocationService{
		locationRepository: locationRepository,
	}
}

func (s *LocationService) GetLocation(id int) (domain.Location, error) {
	location, err := s.locationRepository.GetLocation(id)
	if err != nil {
		return domain.Location{}, err
	}
	weather, err := s.getWeather(location)
	location.Weather = weather
	return location, err
}

func (s *LocationService) AddLocation(location domain.Location) (domain.Location, error) {
	return s.locationRepository.AddLocation(location)
}

func (s *LocationService) GetLocations(userID int) ([]domain.Location, error) {
	locations, err := s.locationRepository.GetLocations(userID)
	if err != nil {
		return []domain.Location{}, err
	}
	for i := range locations {
		weather, err := s.getWeather(locations[i])
		if err != nil {
			return []domain.Location{}, err
		}
		locations[i].Weather = weather
	}
	return locations, err
}

func (s *LocationService) RemoveLocation(id int) error {
	return s.locationRepository.RemoveLocation(id)
}

func (s *LocationService) getWeather(location domain.Location) (domain.Weather, error) {
	apiKey := os.Getenv("API_KEY_WEATHER")
	url := fmt.Sprintf(
		"https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s",
		location.Name,
		apiKey,
	)
	var weather domain.Weather

	resp, err := http.Get(url)
	if err != nil {
		return weather, err
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Printf("response close error: %v", err)
		}
	}()
	if resp.StatusCode != http.StatusOK {
		return weather, fmt.Errorf("weather api error: %s", resp.Status)
	}
	if err := json.NewDecoder(resp.Body).Decode(&weather); err != nil {
		return weather, fmt.Errorf("json decode error")
	}
	return weather, nil
}
