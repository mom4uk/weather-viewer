package unit

import (
	"testing"
	"weather-viewer/internal/services"

	"github.com/stretchr/testify/assert"
)

func TestSearchLocation_success(t *testing.T) {
	fakeRepo := NewFakeRepository()
	fakeWeather := NewFakeWeather()
	service := services.NewLocationService(fakeRepo, fakeWeather)
	location, err := service.GetLocation(1)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 1, location.ID)
	assert.Equal(t, "Москва", location.Name)
	assert.Equal(t, 1, location.ID)
	assert.Equal(t, 55.7522, location.Latitude)
	assert.Equal(t, 37.6156, location.Longitude)
	assert.NotEmpty(t, location.Weather)
}
