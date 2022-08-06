package temphttp

import (
	"log"
	"net/http"
	"os"
	"strings"

	"exercise1/adapters/http/auth"
	fibonacci2 "exercise1/adapters/http/fibonacci"
	"exercise1/adapters/http/math"
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

	mux := mux2.NewRouter()
	mux.Use(authMiddleware.AuthHandler)

	mux.HandleFunc("/add", mathHandler.SumHandler).Methods(http.MethodPost)
	mux.HandleFunc("/fibonacci/{n}", fibHandler.FibHandler).Methods(http.MethodGet)

	return mux
}
