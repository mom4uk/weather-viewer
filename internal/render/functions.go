package render

import (
	"fmt"
	"math"
	"weather-viewer/internal/domain"
)

func weatherIcon(weather domain.Weather) string {
	if len(weather.Weather) == 0 || weather.Weather[0].Icon == "" {
		return "https://openweathermap.org/img/wn/01d@4x.png"
	}

	return fmt.Sprintf("https://openweathermap.org/img/wn/%s@4x.png", weather.Weather[0].Icon)
}

func weatherDescription(weather domain.Weather) string {
	if len(weather.Weather) == 0 {
		return ""
	}

	return weather.Weather[0].Description
}

func temperatureC(kelvin float64) int {
	return int(math.Round(kelvin - 273.15))
}
