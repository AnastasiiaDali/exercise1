package flakiness

import (
	"fmt"
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
			responseStatus  = http.StatusInternalServerError
			flakiness       = r.URL.Query()["flakiness"]
			probability     = 0.0
			randomN         = min + rand.Float64()*(max-min)
			flakinessParams = strings.Split(flakiness[0], ",")
			length          = len(flakinessParams)
		)

		switch length {
		case 3:
			status, err := strconv.Atoi(flakinessParams[1])
			if err != nil {
				log.Fatal("error converting flakinessParams into integer status %w", err)
			}
			responseStatus = status

			sleep := flakinessParams[2]
			parsedDelay, err := time.ParseDuration(sleep)
			fmt.Printf("delay %s", parsedDelay)
			if err != nil {
				log.Fatal("error parsing sleep duration %w", err)
			}
			time.Sleep(parsedDelay)

			w.WriteHeader(responseStatus)
			return
		case 2:
			status, _ := strconv.Atoi(flakinessParams[1])
			responseStatus = status
			w.WriteHeader(responseStatus)
			return
		}

		probability, err := strconv.ParseFloat(strings.Join(flakiness, ""), 32)
		if err != nil {
			log.Fatal("error converting flakiness into an float")
		}

		if randomN <= probability {
			w.WriteHeader(responseStatus)
			return
		}

		handler.ServeHTTP(w, r)
	})
}
