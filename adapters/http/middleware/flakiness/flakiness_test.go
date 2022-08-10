package flakiness

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"exercise1/adapters/http/math"
	mux2 "github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type mockCalculator struct {
	sum int
}

func (c mockCalculator) Add(numbers []int) int {
	return c.sum
}

func TestFlakinessMiddleware(t *testing.T) {
	flakinessMiddlewareTests := []struct {
		Name         string
		Query        string
		ResponseCode int
		SleepTime    int
		Time         string
	}{
		{
			Name:         "Given flakiness 1 should return 500 response",
			Query:        "?flakiness=1",
			ResponseCode: http.StatusInternalServerError,
			SleepTime:    0,
			Time:         "seconds",
		},
		{
			Name:         "Given flakiness 0 should return 200 response",
			Query:        "?num=2&num=3&flakiness=0",
			ResponseCode: http.StatusOK,
			SleepTime:    0,
			Time:         "seconds",
		},
		{
			Name:         "Given flakiness 1 should return 404 response",
			Query:        "?num=2&num=3&flakiness=1,404",
			ResponseCode: http.StatusNotFound,
			SleepTime:    0,
			Time:         "seconds",
		},
		{
			Name:         "Given a delay of 1s response should be delayed for 1s",
			Query:        "?flakiness=1,404,3s",
			ResponseCode: http.StatusNotFound,
			SleepTime:    3,
			Time:         "seconds",
		},
		{
			Name:         "Given a delay of 500ms response should be delayed for 500ms",
			Query:        "?flakiness=1,404,500ms",
			ResponseCode: http.StatusNotFound,
			SleepTime:    500,
			Time:         "milliseconds",
		},
	}

	for _, tc := range flakinessMiddlewareTests {
		t.Run(tc.Name, func(t *testing.T) {

			newMockCalculator := mockCalculator{sum: 8}
			flakinessMiddleware := NewFlakinessMiddleware()
			handler := math.NewHandler(newMockCalculator)

			r := mux2.NewRouter()
			r.Use(flakinessMiddleware.FlakinessHandler)
			r.HandleFunc(fmt.Sprintf("/add"), handler.SumHandler).Methods(http.MethodPost)

			req, err := http.NewRequest(http.MethodPost, "/add"+tc.Query, http.NoBody)
			require.NoError(t, err)
			start := time.Now()
			res := httptest.NewRecorder()
			r.ServeHTTP(res, req)

			actualCode := res.Code
			actualSleepTime := 0

			if tc.Time == "milliseconds" {
				actualSleepTime = int(time.Since(start).Milliseconds())
			} else {
				actualSleepTime = int(time.Since(start).Seconds())
			}

			assert.Equal(t, tc.SleepTime, actualSleepTime)
			assert.Equal(t, tc.ResponseCode, actualCode)
		})
	}
}
