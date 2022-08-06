package fibonacci

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	mux2 "github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type mockFibonacci struct {
	fibNumber string
}

func (f mockFibonacci) FibonacciNumber(number int) (string, error) {
	return f.fibNumber, nil
}

func TestHandler_FibHandler(t *testing.T) {
	newMockFibonacci := mockFibonacci{fibNumber: "610"}

	handler := NewHandler(newMockFibonacci)
	r := mux2.NewRouter()
	r.HandleFunc(fmt.Sprintf("/fibonacci/{n}"), handler.FibHandler).Methods(http.MethodGet)

	req := httptest.NewRequest(http.MethodGet, "/fibonacci/15", http.NoBody)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	resBody, err := io.ReadAll(res.Body)
	assert.NoError(t, err, "failed to read response body")

	var result string
	err = json.Unmarshal(resBody, &result)
	require.NoError(t, err)

	assert.Equal(t, "610", result)

	assert.Equal(t, http.StatusOK, res.Code)
}
