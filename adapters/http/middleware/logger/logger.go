package logger

import (
	"log"
	"net/http"
	"strings"
	"time"
)

type LoggingMiddleware struct {
}

type responseObserver struct {
	http.ResponseWriter
	status      int
	written     int64
	wroteHeader bool
}

func (o *responseObserver) Write(p []byte) (n int, err error) {
	if !o.wroteHeader {
		o.WriteHeader(http.StatusOK)
	}
	n, err = o.ResponseWriter.Write(p)
	o.written += int64(n)
	return
}

func (o *responseObserver) WriteHeader(code int) {
	o.ResponseWriter.WriteHeader(code)
	if o.wroteHeader {
		return
	}
	o.wroteHeader = true
	o.status = code
}

func NewAuthMiddleware() LoggingMiddleware {
	return LoggingMiddleware{}
}

func (m *LoggingMiddleware) AuthHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorizationKey := r.Header.Get("Authorization")
		key := strings.Replace(authorizationKey, "Bearer", "", 1)
		cutKey := key[0:9]
		time := time.Now().Format("02/Jan/2006:15:04:05 -0700")
		o := &responseObserver{ResponseWriter: w}
		handler.ServeHTTP(o, r)

		log.Printf("%s %s %s %s %s %d\n", r.Method, r.URL, r.Body, cutKey, time, o.status)

	})
}
