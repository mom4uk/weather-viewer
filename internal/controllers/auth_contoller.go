package controllers

import (
	"encoding/json"
	"net/http"
	"strings"
	"weather-viewer/internal/apierrors"
	"weather-viewer/internal/dto"
	"weather-viewer/internal/services"
)

type AuthController struct {
	userService    *services.UserService
	sessionService *services.SessionService
	authService    *services.AuthService
}

func NewAuthController(
	userService *services.UserService,
	sessionService *services.SessionService,
	authService *services.AuthService,
) *AuthController {
	return &AuthController{
		userService:    userService,
		sessionService: sessionService,
		authService:    authService,
	}
}

func (c *AuthController) SignUp(w http.ResponseWriter, r *http.Request) {
	login := strings.TrimSpace(r.FormValue("login"))
	password := strings.TrimSpace(r.FormValue("password"))

	if err := dto.ValidateCredentials(login, password); err != nil {
		apierrors.HandleError(w, err)
		return
	}

	ctx := r.Context()
	session, user, err := c.authService.SignUp(ctx, login, password)
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

	if wantsHTML(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	w.WriteHeader(http.StatusCreated)

	res := dto.UserResponse{
		Login: user.Login,
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		apierrors.WriteError(w, "Ошибка при формировании json", http.StatusInternalServerError)
	}
}

func (c *AuthController) SignIn(w http.ResponseWriter, r *http.Request) {
	login := strings.TrimSpace(r.FormValue("login"))
	password := strings.TrimSpace(r.FormValue("password"))

	if err := dto.ValidateCredentials(login, password); err != nil {
		apierrors.HandleError(w, err)
		return
	}
	ctx := r.Context()
	session, err := c.authService.SignIn(ctx, login, password)
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

	if wantsHTML(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (c *AuthController) SignOut(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_token")
	if err != nil {
		if wantsHTML(r) {
			http.Redirect(w, r, "/sign-in", http.StatusSeeOther)
			return
		}
		w.WriteHeader(http.StatusNoContent)
		return
	}
	// я хз, как мне правильно обрабатывать logout. Нужно ли мне обрабатывать как-то ошибки или все таки, главное,
	// что я просто очистил куки, а очистились они в бд или она с ошибкой упала мне не важно?
	ctx := r.Context()
	err = c.authService.SignOut(ctx, cookie.Value)
	if err != nil {
		if wantsHTML(r) {
			http.Redirect(w, r, "/sign-in", http.StatusSeeOther)
			return
		}
		w.WriteHeader(http.StatusNoContent)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:   "session_token",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	})

	if wantsHTML(r) {
		http.Redirect(w, r, "/sign-in", http.StatusSeeOther)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func wantsHTML(r *http.Request) bool {
	return strings.Contains(r.Header.Get("Accept"), "text/html")
}
