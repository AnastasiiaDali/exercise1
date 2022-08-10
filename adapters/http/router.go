package temphttp

import (
	"log"
	"net/http"
	"os"
	"strings"

	fibonacci2 "exercise1/adapters/http/fibonacci"
	"exercise1/adapters/http/math"
	"exercise1/adapters/http/middleware/auth"
	"exercise1/adapters/http/middleware/flakiness"
	logger2 "exercise1/adapters/http/middleware/logger"
	mux2 "github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func NewRouter(calculator Calculator, fibonacci Fibonacci) http.Handler {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	keys := os.Getenv("AUTH_KEYS")
	authKeys := strings.Split(keys, ",")

	var authorisedUsers []string
	for _, key := range authKeys {
		authorisedUsers = append(authorisedUsers, key)
	}

	mathHandler := math.NewHandler(calculator)
	fibHandler := fibonacci2.NewHandler(fibonacci)

	authMiddleware := auth.NewAuthMiddleware(authorisedUsers)
	loggingMiddleware := logger2.NewAuthMiddleware()
	flakinessMiddleware := flakiness.NewFlakinessMiddleware()

	mux := mux2.NewRouter()
	mux.Use(loggingMiddleware.AuthHandler, authMiddleware.AuthHandler, flakinessMiddleware.FlakinessHandler)

	mux.HandleFunc("/add", mathHandler.SumHandler).Methods(http.MethodPost)
	mux.HandleFunc("/fibonacci/{n}", fibHandler.FibHandler).Methods(http.MethodGet)

	return mux
}
