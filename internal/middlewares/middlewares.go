package middlewares

import (
	"context"
	"net/http"
	"weather-viewer/internal/apierrors"
	"weather-viewer/internal/services"
)

type Middleware func(http.Handler) http.Handler

func JSON() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			next.ServeHTTP(w, r)
		})
	}
}

func Auth(s *services.SessionService) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cookie, err := r.Cookie("session_token")
			if err != nil {
				apierrors.WriteError(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			ctx := r.Context()
			userID, err := s.GetUserID(ctx, cookie.Value)
			if err != nil {
				apierrors.WriteError(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
			ctx = context.WithValue(ctx, "user_id", userID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func Chain(mw ...Middleware) Middleware {
	return func(final http.Handler) http.Handler {
		for i := len(mw) - 1; i >= 0; i-- {
			final = mw[i](final)
		}
		return final
	}
}
