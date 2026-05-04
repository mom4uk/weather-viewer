package services

import (
	"weather-viewer/internal/domain"
	"weather-viewer/internal/utilities"
)

type AuthService struct {
	sessionService *SessionService
	userService    *UserService
}

func NewAuthService(sessionService *SessionService, userService *UserService) *AuthService {
	return &AuthService{
		sessionService: sessionService,
		userService:    userService,
	}
}

func (s *AuthService) RegisterUser(login, password string) (domain.Session, domain.User, error) {
	user, err := s.userService.CreateUser(login, password)
	if err != nil {
		return domain.Session{}, domain.User{}, err
	}

	session, err := s.sessionService.CreateSession(user.ID)
	if err != nil {
		return domain.Session{}, domain.User{}, err
	}
	return session, user, nil
}

func (s *AuthService) LoginUser(login, password string) (domain.Session, error) {
	user, err := s.userService.GetUserByLogin(login)
	if err != nil {
		return domain.Session{}, domain.ErrIncorrectCredentials
	}

	if !utilities.ComparePasswords(user.Password, password) {
		return domain.Session{}, domain.ErrIncorrectCredentials
	}

	session, err := s.sessionService.CreateSession(user.ID)
	if err != nil {
		return domain.Session{}, err
	}
	return session, nil
}
