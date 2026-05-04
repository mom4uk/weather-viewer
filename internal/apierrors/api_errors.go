package apierrors

import (
	"encoding/json"
	"errors"
	"net/http"
	"weather-viewer/internal/domain"
)

func WriteError(w http.ResponseWriter, message string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	_ = json.NewEncoder(w).Encode(domain.ErrorResponse{
		Message: message,
	})
}

func HandleError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, domain.ErrInvalidId):
		WriteError(w, "Некорректное значение в id", http.StatusBadRequest)
		return
	case errors.Is(err, domain.ErrLocationNotFound):
		WriteError(w, "Данная локация не найдена", http.StatusNotFound)
		return
	case errors.Is(err, domain.ErrInvalidName):
		WriteError(w, "Некорректное значение в name", http.StatusBadRequest)
		return
	case errors.Is(err, domain.ErrInvalidLatitude):
		WriteError(w, "Некорректное значение в latitude", http.StatusBadRequest)
		return
	case errors.Is(err, domain.ErrInvalidLongitude):
		WriteError(w, "Некорректное значение в longitude", http.StatusBadRequest)
		return
	case errors.Is(err, domain.ErrLocationAlreadyExists):
		WriteError(w, "Такая локация уже существует", http.StatusConflict)
		return
	case errors.Is(err, domain.ErrUserNotFound):
		WriteError(w, "Пользователь не найден", http.StatusNotFound)
		return
	case errors.Is(err, domain.ErrLoginInvalidLength):
		WriteError(w, "Длина логина должен быть от 6 до 20 символов", http.StatusUnprocessableEntity)
		return
	case errors.Is(err, domain.ErrPasswordInvalidLength):
		WriteError(w, "Длина пароля должна быть от 6 до 20 символов", http.StatusUnprocessableEntity)
		return
	case errors.Is(err, domain.ErrInvalidLogin):
		WriteError(w, "Логин может состоять только из латинских букв и цифр", http.StatusUnprocessableEntity)
		return
	case errors.Is(err, domain.ErrUserAlreadyExists):
		WriteError(w, "Пользователь с таким логином уже существует", http.StatusConflict)
		return
	case errors.Is(err, domain.ErrAbsenceOfLoginPass):
		WriteError(w, "Не передан логин и/или пароль", http.StatusBadRequest)
		return
	case errors.Is(err, domain.ErrIncorrectCredentials):
		WriteError(w, "Неверный логин или пароль", http.StatusUnauthorized)
		return
	default:
		WriteError(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
