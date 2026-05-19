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

func Response(err error) (string, int) {
	var weatherErr domain.WeatherAPIError
	switch {
	case errors.Is(err, domain.ErrInvalidId):
		return "Некорректное значение в id", http.StatusBadRequest
	case errors.Is(err, domain.ErrLocationNotFound):
		return "Данная локация не найдена", http.StatusNotFound
	case errors.Is(err, domain.ErrInvalidName):
		return "Некорректное значение в name", http.StatusBadRequest
	case errors.Is(err, domain.ErrInvalidLatitude):
		return "Некорректное значение в latitude", http.StatusBadRequest
	case errors.Is(err, domain.ErrInvalidLongitude):
		return "Некорректное значение в longitude", http.StatusBadRequest
	case errors.Is(err, domain.ErrLocationAlreadyExists):
		return "Такая локация уже существует", http.StatusConflict
	case errors.Is(err, domain.ErrUserNotFound):
		return "Пользователь не найден", http.StatusNotFound
	case errors.Is(err, domain.ErrLoginInvalidLength):
		return "Длина логина должна быть от 6 до 20 символов", http.StatusUnprocessableEntity
	case errors.Is(err, domain.ErrPasswordInvalidLength):
		return "Длина пароля должна быть от 6 до 20 символов", http.StatusUnprocessableEntity
	case errors.Is(err, domain.ErrInvalidLogin):
		return "Логин может состоять только из латинских букв и цифр", http.StatusUnprocessableEntity
	case errors.Is(err, domain.ErrUserAlreadyExists):
		return "Пользователь с таким логином уже существует", http.StatusConflict
	case errors.Is(err, domain.ErrAbsenceOfLoginPass):
		return "Не передан логин и/или пароль", http.StatusBadRequest
	case errors.Is(err, domain.ErrIncorrectCredentials):
		return "Неверный логин или пароль", http.StatusUnauthorized
	case errors.Is(err, domain.ErrPasswordsNotMatch):
		return "Пароли не совпадают", http.StatusBadRequest
	case errors.As(err, &weatherErr):
		return weatherErr.Message, http.StatusBadGateway
	default:
		return err.Error(), http.StatusInternalServerError
	}
}

func HandleError(w http.ResponseWriter, err error) {
	message, code := Response(err)
	WriteError(w, message, code)
}
