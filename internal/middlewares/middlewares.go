package middlewares

import (
	"context"
	"net/http"
	"weather-viewer/internal/apierrors"
	"weather-viewer/internal/services"
)

func JSON(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func Auth(s *services.SessionService) func(handler http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cookie, err := r.Cookie("session_token")
			if err != nil {
				apierrors.WriteError(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			userID, err := s.Authenticate(cookie.Value)
			if err != nil {
				apierrors.WriteError(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), "user_id", userID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
