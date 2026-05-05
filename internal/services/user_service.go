package services

import (
	"weather-viewer/internal/domain"
	"weather-viewer/internal/repositories"
)

type UserService struct {
	repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) CreateUser(login, hash string) (domain.User, error) {
	return s.repo.CreateUser(login, hash)
}

func (s *UserService) GetUserByLogin(login string) (domain.User, error) {
	return s.repo.GetUserByLogin(login)
}
