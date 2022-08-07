package logger

import (
	"log"
	"net/http"
)

type LoggingMiddleware struct {
}

func NewAuthMiddleware() LoggingMiddleware {
	return LoggingMiddleware{}
}

func (m *LoggingMiddleware) AuthHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.Method, r.URL, r.Body)
		handler.ServeHTTP(w, r)
	})
}
