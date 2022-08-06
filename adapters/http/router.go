package temphttp

import (
	"log"
	"net/http"
	"os"
	"strings"

	"exercise1/adapters/http/auth"
	"exercise1/adapters/http/math"
	mux2 "github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func NewRouter(calculator Calculator) http.Handler {
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

	handler := math.NewHandler(calculator)

	authMiddleware := auth.NewAuthMiddleware(authorisedUsers)

	mux := mux2.NewRouter()
	mux.Use(authMiddleware.AuthHandler)

	mux.HandleFunc("/add", handler.SumHandler).Methods(http.MethodPost)

	return mux
}
