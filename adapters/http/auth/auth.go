package auth

import (
	"fmt"
	"net/http"
	"strings"
)

const bearerPrefix = "Bearer "

type AuthorisedUsers []string

type AuthMiddleware struct {
	users AuthorisedUsers
}

func NewAuthMiddleware(au AuthorisedUsers) AuthMiddleware {
	return AuthMiddleware{users: au}
}

func (m *AuthMiddleware) AuthHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bearerToken := strings.TrimSpace(r.Header.Get("Authorization"))
		fmt.Printf("users are here %s", m.users)
		if !strings.HasPrefix(bearerToken, bearerPrefix) {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		token := strings.TrimPrefix(bearerToken, bearerPrefix)

		isContain := contains(m.users, token)
		if !isContain {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		handler.ServeHTTP(w, r)
	})
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
