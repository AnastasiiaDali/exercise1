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
		time := time.Now().Format("2006-02-01T15:04:05Z")
		o := &responseObserver{ResponseWriter: w}
		handler.ServeHTTP(o, r)

		log.Printf("Time:%s;\n Method: %s;\n URL: %s;\n Auth Token%s;\n Status Code: %d;\n ContentLength: %d\n", time, r.Method, r.URL, cutKey, o.status, r.ContentLength)
	})
}
