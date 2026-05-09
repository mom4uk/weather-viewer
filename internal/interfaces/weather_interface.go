package interfaces

import "weather-viewer/internal/domain"

type Weather interface {
	GetWeather(location domain.Location) (domain.Weather, error)
}
