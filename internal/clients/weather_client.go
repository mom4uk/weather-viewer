package clients

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"weather-viewer/internal/domain"
)

type WeatherClient struct {
	URL    string
	APIKey string
	Client *http.Client
}

func NewWeatherClient(url, apiKey string, client *http.Client) *WeatherClient {
	return &WeatherClient{
		URL:    url,
		APIKey: apiKey,
		Client: client,
	}
}

func (c *WeatherClient) GetWeather(location domain.Location) (domain.Weather, error) {
	url := fmt.Sprintf(
		"%s/data/2.5/weather?q=%s&appid=%s",
		c.URL,
		location.Name,
		c.APIKey,
	)

	var weather domain.Weather

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return domain.Weather{}, err
	}

	resp, err := c.Client.Do(req)
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
