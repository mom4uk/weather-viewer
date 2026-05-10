package middlewares

import (
	"context"
	"net/http"
	"time"
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

			session, err := s.GetSession(cookie.Value)
			if err != nil {
				apierrors.WriteError(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			if session.IsExpired(time.Now()) {
				apierrors.WriteError(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), "session", session)
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
