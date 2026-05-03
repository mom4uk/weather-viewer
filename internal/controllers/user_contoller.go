package controllers

import (
	"encoding/json"
	"net/http"
	"strings"
	"weather-viewer/internal/apierrors"
	"weather-viewer/internal/domain"
	"weather-viewer/internal/dto"
	"weather-viewer/internal/services"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{
		userService: userService,
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

	if err := json.NewEncoder(w).Encode(res); err != nil {
		apierrors.WriteError(w, "Ошибка при формировании json", http.StatusInternalServerError)
	}
}
