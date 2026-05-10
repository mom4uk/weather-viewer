package fixtures

import "weather-viewer/internal/domain"

func GetMoscowWeather() domain.Weather {
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
				ID:          804,
				Main:        "Clouds",
				Description: "overcast clouds",
				Icon:        "05d",
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
			Temp:     285.85,
			Pressure: 1019,
			Humidity: 44,
			TempMin:  285.01,
			TempMax:  286.39,
		},

		Visibility: 10000,

		Wind: struct {
			Speed float64 `json:"speed"`
			Deg   int     `json:"deg"`
		}{
			Speed: 4.62,
			Deg:   138,
		},

		Clouds: struct {
			All int `json:"all"`
		}{
			All: 100,
		},

		Dt: 1778422833,

		Sys: struct {
			Type    int     `json:"type"`
			ID      int     `json:"id"`
			Message float64 `json:"message"`
			Country string  `json:"country"`
			Sunrise int64   `json:"sunrise"`
			Sunset  int64   `json:"sunset"`
		}{
			Type:    2,
			ID:      2094500,
			Message: 0,
			Country: "RU",
			Sunrise: 1778376482,
			Sunset:  1778433835,
		},

		ID:   524901,
		Name: "Moscow",
		Cod:  200,
	}
}

func GetSpbWeather() domain.Weather {
	return domain.Weather{
		Coord: struct {
			Lon float64 `json:"lon"`
			Lat float64 `json:"lat"`
		}{
			Lon: 30.2642,
			Lat: 59.8944,
		},

		Weather: []struct {
			ID          int    `json:"id"`
			Main        string `json:"main"`
			Description string `json:"description"`
			Icon        string `json:"icon"`
		}{
			{
				ID:          800,
				Main:        "Clear",
				Description: "clear sky",
				Icon:        "01d",
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
			Temp:     288.23,
			Pressure: 1014,
			Humidity: 33,
			TempMin:  288.23,
			TempMax:  288.23,
		},

		Visibility: 10000,

		Wind: struct {
			Speed float64 `json:"speed"`
			Deg   int     `json:"deg"`
		}{
			Speed: 7,
			Deg:   150,
		},

		Clouds: struct {
			All int `json:"all"`
		}{
			All: 0,
		},

		Dt: 1778426404,

		Sys: struct {
			Type    int     `json:"type"`
			ID      int     `json:"id"`
			Message float64 `json:"message"`
			Country string  `json:"country"`
			Sunrise int64   `json:"sunrise"`
			Sunset  int64   `json:"sunset"`
		}{
			Type:    1,
			ID:      8926,
			Message: 0,
			Country: "RU",
			Sunrise: 1778376870,
			Sunset:  1778436976,
		},

		ID:   498817,
		Name: "Saint Petersburg",
		Cod:  200,
	}
}
