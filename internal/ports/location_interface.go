package ports

import "weather-viewer/internal/domain"

type LocationRepository interface {
	GetLocation(id int) (domain.Location, error)
}
