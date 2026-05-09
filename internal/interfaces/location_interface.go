package interfaces

import "weather-viewer/internal/domain"

type LocationRepository interface {
	GetLocation(id int) (domain.Location, error)
	AddLocation(location domain.Location) (domain.Location, error)
	GetLocations(sessionToken int) ([]domain.Location, error)
	RemoveLocation(id int) error
}
