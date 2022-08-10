package flakiness

import (
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
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
		var (
			responseStatus = http.StatusInternalServerError
			flakiness      = r.URL.Query()["flakiness"]
			probability    = 0.0
		)

		random := min + rand.Float64()*(max-min)

		flakinessParams := strings.Split(flakiness[0], ",")

		if len(flakinessParams) >= 3 {
			status, _ := strconv.Atoi(flakinessParams[1])
			responseStatus = status
			sleep := flakinessParams[2]
			parsedDelay, err := time.ParseDuration(sleep)
			if err != nil {
				log.Fatal("error parsing sleep duration %w", err)
			}

			time.Sleep(parsedDelay)
			w.WriteHeader(responseStatus)
			return
		} else if len(flakinessParams) >= 2 {
			status, _ := strconv.Atoi(flakinessParams[1])
			responseStatus = status
			w.WriteHeader(responseStatus)
			return
		} else {
			f := strings.Join(flakiness, "")
			p, err := strconv.ParseFloat(f, 32)
			if err != nil {
				log.Fatal("error converting flakiness into an float")
			}

			probability = p
		}

		if random <= probability {
			w.WriteHeader(responseStatus)
			return
		}

		handler.ServeHTTP(w, r)
	})
}
