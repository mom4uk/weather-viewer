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
	case errors.Is(err, domain.ErrIncorrectId):
		WriteError(w, "Некорректное значение в id", http.StatusBadRequest)
		return
	default:
		WriteError(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
