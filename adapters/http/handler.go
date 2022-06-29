package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"exercise1/domain/calculator"
	"exercise1/domain/formatter"
	"exercise1/helpers/string_to_int_converter"
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
	values = string_to_int_converter.StringToIntConverter(numbers)

	//pass number to add function and get the sum and then format it
	sum := calculator.Add(values)
	formattedSum := formatter.FormatNumber(sum)

	res, err := json.Marshal(formattedSum)

	w.Header().Add("Content-type", "text/plain")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprint(w, "Healthy")
}
