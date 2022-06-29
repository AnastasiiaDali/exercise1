package http_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	http2 "exercise1/adapters/http"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestMathHandler(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/add", http2.MathHandler).Methods(http.MethodPost)

	req := httptest.NewRequest(http.MethodPost, "/add?num=50&num=50", http.NoBody)
	res := httptest.NewRecorder()

	http2.MathHandler(res, req)

	var sum string
	expected := "100"

	err := json.Unmarshal(res.Body.Bytes(), &sum)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, expected, sum)
}
