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
	hash, err := utilities.HashPassword(password)
	if err != nil {
		return domain.Session{}, domain.User{}, err
	}
	user, err := s.userService.CreateUser(login, hash)
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

	if err := utilities.ComparePasswords(user.Password, password); err != nil {
		return domain.Session{}, domain.ErrIncorrectCredentials
	}

	session, err := s.sessionService.CreateSession(user.ID)
	if err != nil {
		return domain.Session{}, err
	}
	return session, nil
}

func (s *AuthService) LogoutUser(sessionID string) error {
	err := s.sessionService.DeleteSession(sessionID)
	if err != nil {
		return err
	}
	return nil
}
