package unit

import (
	"weather-viewer/internal/domain"
)

type Weather struct {
}

func NewFakeWeather() *Weather {
	return &Weather{}
}

func (w *Weather) GetWeather(location domain.Location) (domain.Weather, error) {
	return domain.Weather{
		Coord: struct {
			Lon float64 `json:"lon"`
			Lat float64 `json:"lat"`
		}{
			Lon: 37.6156,
			Lat: 55.7522,
		},

		Weather: []struct {
			ID          int    `json:"id"`
			Main        string `json:"main"`
			Description string `json:"description"`
			Icon        string `json:"icon"`
		}{
			{
				ID:          501,
				Main:        "Rain",
				Description: "moderate rain",
				Icon:        "10d",
			},
		},

		Base: "stations",

		Main: struct {
			Temp     float64 `json:"temp"`
			Pressure int     `json:"pressure"`
			Humidity int     `json:"humidity"`
			TempMin  float64 `json:"temp_min"`
			TempMax  float64 `json:"temp_max"`
		}{
			Temp:     295.39,
			Pressure: 1009,
			Humidity: 49,
			TempMin:  295.39,
			TempMax:  298.44,
		},

		Visibility: 10000,

		Wind: struct {
			Speed float64 `json:"speed"`
			Deg   int     `json:"deg"`
		}{
			Speed: 1.94,
			Deg:   44,
		},

		Clouds: struct {
			All int `json:"all"`
		}{
			All: 69,
		},

		Dt: 1778165684,

		Sys: struct {
			Type    int     `json:"type"`
			ID      int     `json:"id"`
			Message float64 `json:"message"`
			Country string  `json:"country"`
			Sunrise int64   `json:"sunrise"`
			Sunset  int64   `json:"sunset"`
		}{
			Type:    2,
			ID:      2018597,
			Message: 0,
			Country: "RU",
			Sunrise: 1778111800,
			Sunset:  1778167900,
		},

		ID:   524901,
		Name: "Moscow",
		Cod:  200,
	}, nil
}
