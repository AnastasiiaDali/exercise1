package flakiness

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

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
	}{
		{
			Name:         "Given flakiness 1 should return 500 response",
			Query:        "?flakiness=1",
			ResponseCode: http.StatusInternalServerError,
		},
		{
			Name:         "Given flakiness 0 should return 200 response",
			Query:        "?num=2&num=3&flakiness=0",
			ResponseCode: http.StatusOK,
		},
		{
			Name:         "Given flakiness 1 should return 404 response",
			Query:        "?num=2&num=3&flakiness=1,404",
			ResponseCode: http.StatusOK,
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

			res := httptest.NewRecorder()
			r.ServeHTTP(res, req)

			gotCode := res.Code

			assert.Equal(t, tc.ResponseCode, gotCode)
		})
	}
}
