package repositories

import (
	"database/sql"
	"errors"
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

func (r *UserRepository) CreateUser(login, password string) (domain.User, error) {
	query := "INSERT INTO users (login, password) VALUES ($1, $2) RETURNING id"
	user := domain.User{
		Login:    login,
		Password: password,
	}
	err := r.db.QueryRow(query, login, password).Scan(&user.ID)

	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return domain.User{}, domain.ErrUserAlreadyExists
		}
		return domain.User{}, err
	}

	return user, nil
}

func (r *UserRepository) GetUserByLogin(login string) (domain.User, error) {
	query := "SELECT id, login, password FROM users WHERE login = $1"

	var user domain.User
	err := r.db.QueryRow(query, login).Scan(&user.ID, &user.Login, &user.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.User{}, domain.ErrUserNotFound
		}
		return domain.User{}, err
	}

	return user, nil
}
