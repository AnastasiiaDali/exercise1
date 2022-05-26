package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"exercise1/domain"
	"exercise1/helpers"
	"github.com/gorilla/mux"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	var numbers []string
	var values []int

	num := r.URL.Query()

	for k, v := range num {
		numbers = v
		fmt.Println(k, " => ", v)
	}

	values = helpers.StringToIntConverter(numbers)

	sum := domain.Add(values)

	res, err := json.Marshal(sum)

	w.Header().Add("Content-type", "text/plain")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/add", handlerFunc).Methods(http.MethodPost)
	fmt.Printf("Starting server for testing HTTP POST...\n")
	http.ListenAndServe(":8081", r)
}
