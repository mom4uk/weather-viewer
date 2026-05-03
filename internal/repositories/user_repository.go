package repositories

import (
	"database/sql"
	"fmt"
	"strings"
	"weather-viewer/internal/domain"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) CreateUser(user domain.User) (domain.User, error) {
	query := "INSERT INTO users (login, password) VALUES ($1, $2) RETURNING id"
	err := r.db.QueryRow(query, user.Login, user.Password).Scan(&user.ID)
	fmt.Print(user)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return domain.User{}, domain.ErrUserAlreadyExists
		}
		return domain.User{}, err
	}

	return user, nil
}
