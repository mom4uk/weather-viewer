package services

import (
	"context"
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

func (s *AuthService) RegisterUser(ctx context.Context, login, password string) (domain.Session, domain.User, error) {
	hash, err := utilities.HashPassword(password)
	if err != nil {
		return domain.Session{}, domain.User{}, err
	}
	user, err := s.userService.CreateUser(login, hash)
	if err != nil {
		return domain.Session{}, domain.User{}, err
	}

	session, err := s.sessionService.CreateSession(ctx, user.ID)
	if err != nil {
		return domain.Session{}, domain.User{}, err
	}
	return session, user, nil
}

func (s *AuthService) LoginUser(ctx context.Context, login, password string) (domain.Session, error) {
	user, err := s.userService.GetUserByLogin(login)
	if err != nil {
		return domain.Session{}, domain.ErrIncorrectCredentials
	}

	if err := utilities.ComparePasswords(user.Password, password); err != nil {
		return domain.Session{}, domain.ErrIncorrectCredentials
	}

	session, err := s.sessionService.CreateSession(ctx, user.ID)
	if err != nil {
		return domain.Session{}, err
	}
	return session, nil
}

func (s *AuthService) LogoutUser(ctx context.Context, sessionID string) error {
	err := s.sessionService.DeleteSession(ctx, sessionID)
	if err != nil {
		return err
	}
	return nil
}
