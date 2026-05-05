package controllers

import (
	"encoding/json"
	"net/http"
	"strings"
	"weather-viewer/internal/apierrors"
	"weather-viewer/internal/dto"
	"weather-viewer/internal/services"
)

type UserController struct {
	userService    *services.UserService
	sessionService *services.SessionService
	authService    *services.AuthService
}

func NewUserController(
	userService *services.UserService,
	sessionService *services.SessionService,
	authService *services.AuthService,
) *UserController {
	return &UserController{
		userService:    userService,
		sessionService: sessionService,
		authService:    authService,
	}
}

func (c *UserController) RegisterUser(w http.ResponseWriter, r *http.Request) {
	login := strings.TrimSpace(r.FormValue("login"))
	password := strings.TrimSpace(r.FormValue("password"))

	if err := dto.ValidateCredentials(login, password); err != nil {
		apierrors.HandleError(w, err)
		return
	}

	session, user, err := c.authService.RegisterUser(login, password)
	if err != nil {
		apierrors.HandleError(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	res := dto.UserResponse{
		Login: user.Login,
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

func (c *UserController) LoginUser(w http.ResponseWriter, r *http.Request) {
	login := strings.TrimSpace(r.FormValue("login"))
	password := strings.TrimSpace(r.FormValue("password"))

	if err := dto.ValidateCredentials(login, password); err != nil {
		apierrors.HandleError(w, err)
		return
	}

	session, err := c.authService.LoginUser(login, password)
	if err != nil {
		apierrors.HandleError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    session.ID,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
	})
}

func (c *UserController) LogoutUser(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_token")
	if err != nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	err = c.authService.LogoutUser(cookie.Value)
	if err != nil {
		apierrors.HandleError(w, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)

	http.SetCookie(w, &http.Cookie{
		Name:   "session_token",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	})
}
