package http

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func MathServer() {
	r := mux.NewRouter()
	r.HandleFunc("/add", MathHandler).Methods(http.MethodPost)
	fmt.Printf("Starting server on port 8081...\n")
	http.ListenAndServe(":8081", r)
}
