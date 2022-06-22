package http

import (
	"encoding/json"
	"net/http"

	"exercise1/domain"
	"exercise1/helpers"
)

func MathHandler(w http.ResponseWriter, r *http.Request) {
	var numbers []string
	var values []int

	//get numbers from url
	num := r.URL.Query()
	for _, v := range num {
		numbers = v
	}

	//convert number into array of integers
	values = helpers.StringToIntConverter(numbers)

	//pass number to add function and get the sum
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
