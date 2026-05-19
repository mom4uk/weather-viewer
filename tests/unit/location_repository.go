package unit

import (
	"weather-viewer/internal/domain"
)

type FakeRepository struct {
}

func NewFakeRepository() *FakeRepository {
	return &FakeRepository{}
}

func (repo *FakeRepository) GetLocation(locationID int) (domain.Location, error) {
	return domain.Location{
		ID: 1, Name: "Москва", UserID: 1, Latitude: 0, Longitude: 0,
	}, nil
}

func (repo *FakeRepository) AddLocation(location domain.Location) (domain.Location, error) {
	return domain.Location{}, nil
}

func (repo *FakeRepository) GetLocations(sessionToken int) ([]domain.Location, error) {
	return []domain.Location{
		{ID: 1, Name: "Москва", UserID: 1, Latitude: 0, Longitude: 0},
	}, nil
}

func (repo *FakeRepository) RemoveLocation(id int) error {
	return nil
}
