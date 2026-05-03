package controllers

import (
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
	"strings"
	"time"
	"weather-viewer/internal/apierrors"
	"weather-viewer/internal/domain"
	"weather-viewer/internal/dto"
	"weather-viewer/internal/services"
)

type UserController struct {
	userService    *services.UserService
	sessionService *services.SessionService
}

func NewUserController(userService *services.UserService, sessionService *services.SessionService) *UserController {
	return &UserController{
		userService:    userService,
		sessionService: sessionService,
	}
}

func (c *UserController) RegisterUser(w http.ResponseWriter, r *http.Request) {
	login := strings.TrimSpace(r.FormValue("login"))
	password := strings.TrimSpace(r.FormValue("password"))

	if err := dto.ValidateCredentials(login, password); err != nil {
		apierrors.HandleError(w, err)
		return
	}

	user := domain.User{
		Login:    login,
		Password: password,
	}

	result, err := c.userService.RegisterUser(user)
	if err != nil {
		apierrors.HandleError(w, err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	res := dto.UserResponse{
		Login: result.Login,
	}

	session := domain.Session{
		ID:        uuid.New().String(),
		UserID:    result.ID,
		ExpiresAt: time.Now().Add(1 * time.Hour),
	}

	err = c.sessionService.CreateSession(session)
	if err != nil {
		apierrors.HandleError(w, err)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    session.ID,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
	})

	if err := json.NewEncoder(w).Encode(res); err != nil {
		apierrors.WriteError(w, "Ошибка при формировании json", http.StatusInternalServerError)
	}
}
