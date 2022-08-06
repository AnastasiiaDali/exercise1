package math

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

type mockCalculator struct {
	sum int
}

func (c mockCalculator) Add(numbers []int) int {
	return c.sum
}

func TestHandler_SumHandler(t *testing.T) {
	newMockCalculator := mockCalculator{sum: 8}

	handler := NewHandler(newMockCalculator)
	r := mux2.NewRouter()
	r.HandleFunc(fmt.Sprintf("/add"), handler.SumHandler).Methods(http.MethodPost)

	req := httptest.NewRequest(http.MethodPost, "/add?num=4&num=4", http.NoBody)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	resBody, err := io.ReadAll(res.Body)
	assert.NoError(t, err, "failed to read response body")

	var result string
	err = json.Unmarshal(resBody, &result)
	require.NoError(t, err)

	assert.Equal(t, "8", result)

	assert.Equal(t, http.StatusOK, res.Code)
}
