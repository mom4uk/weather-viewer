package services

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"weather-viewer/internal/domain"
)

type WeatherService struct {
}

func NewWeatherService() *WeatherService {
	return &WeatherService{}
}

func (s *WeatherService) GetWeather(location domain.Location) (domain.Weather, error) {
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
