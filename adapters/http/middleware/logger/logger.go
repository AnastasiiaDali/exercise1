package logger

import (
	"log"
	"net/http"
	"strings"
)

type LoggingMiddleware struct {
}

func NewAuthMiddleware() LoggingMiddleware {
	return LoggingMiddleware{}
}

func (m *LoggingMiddleware) AuthHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorizationKey := r.Header.Get("Authorization")
		key := strings.Replace(authorizationKey, "Bearer", "", 1)
		cutKey := key[0:9]
		log.Printf("%s %s %s %s\n", r.Method, r.URL, r.Body, cutKey)
		handler.ServeHTTP(w, r)
	})
}
