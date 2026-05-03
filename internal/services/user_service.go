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

func (s *UserService) RegisterUser(user domain.User) (domain.User, error) {
	return s.repo.CreateUser(user)
}
