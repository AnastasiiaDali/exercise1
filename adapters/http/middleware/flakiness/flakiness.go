package flakiness

import (
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
)

const (
	min = 0.0
	max = 1.0
)

type FlakinessMiddleware struct {
}

func NewFlakinessMiddleware() FlakinessMiddleware {
	return FlakinessMiddleware{}
}

func (f *FlakinessMiddleware) FlakinessHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		random := min + rand.Float64()*(max-min)

		flakiness := r.URL.Query()["flakiness"]

		if len(flakiness) == 0 {
			handler.ServeHTTP(w, r)
			return
		}

		f := strings.Join(flakiness, "")
		probability, err := strconv.ParseFloat(f, 32)
		if err != nil {
			log.Fatal("error converting flakiness into an float")
		}

		if random <= probability {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		handler.ServeHTTP(w, r)
	})
}
