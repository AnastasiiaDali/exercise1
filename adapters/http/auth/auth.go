package auth

import (
	"net/http"
	"strings"
)

//Authorization: Bearer SUPER_SECRET_API_KEY

const bearerPrefix = "Bearer "

type AuthMiddleware struct {
}

func NewAuthMiddleware() AuthMiddleware {
	return AuthMiddleware{}
}

func (m *AuthMiddleware) AuthHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bearerToken := strings.TrimSpace(r.Header.Get("Authorization"))

		if !strings.HasPrefix(bearerToken, bearerPrefix) {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		token := strings.TrimPrefix(bearerToken, bearerPrefix)

		if token != "SUPER_SECRET_API_KEY" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		handler.ServeHTTP(w, r)
	})
}
