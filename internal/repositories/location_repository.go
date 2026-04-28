package repositories

import (
	"database/sql"
	"errors"
	"weather-viewer/internal/domain"
)

type LocationRepository struct {
	db *sql.DB
}

func NewLocationRepository(db *sql.DB) *LocationRepository {
	return &LocationRepository{
		db: db,
	}
}
func (r *LocationRepository) GetLocation(id int) (domain.Location, error) {
	query := `SELECT id, name, user_id, latitude, longitude FROM locations WHERE id = $1`

	var location domain.Location
	err := r.db.QueryRow(query, id).Scan(
		&location.ID,
		&location.Name,
		&location.UserID,
		&location.Latitude,
		&location.Longitude,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Location{}, domain.ErrIncorrectNotFound
		}
		return domain.Location{}, err
	}

	return location, nil
}
