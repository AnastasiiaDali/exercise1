package fibonacci

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type TempFibonacci interface {
	FibonacciNumber(numbers int) (string, error)
}

type Handler struct {
	fibonacci TempFibonacci
}

func NewHandler(fibonacci TempFibonacci) *Handler {
	return &Handler{fibonacci: fibonacci}
}

func (h *Handler) FibHandler(w http.ResponseWriter, r *http.Request) {
	n := mux.Vars(r)["n"]

	//convert number into integer
	number, err := strconv.Atoi(n)
	if err != nil {
		panic(err)
	}

	fibonacciNumber, err := h.fibonacci.FibonacciNumber(number)
	if err != nil {
		panic(err)
	}

	res, err := json.Marshal(fibonacciNumber)

	w.Header().Add("Content-type", "text/plain")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
