package unit

import (
	"weather-viewer/internal/domain"
	"weather-viewer/tests/fixtures"
)

type Weather struct {
}

func NewFakeWeather() *Weather {
	return &Weather{}
}

func (w *Weather) GetWeather(location domain.Location) (domain.Weather, error) {
	return fixtures.GetMoscowWeather(), nil
}
