package auth

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	test_helpers "exercise1/test-helpers"
	"github.com/stretchr/testify/assert"
)

func TestNewAuthMiddleware(t *testing.T) {
	const testToken = "testtoken"

	authorisedUsers := []string{testToken}

	testCases := []struct {
		name               string
		request            *http.Request
		expectedStatusCode int
		isAuthorized       bool
	}{
		{name: "continue request with authorized user", request: newRequestWithAuthHeader(t, fmt.Sprintln("Bearer", testToken)), expectedStatusCode: http.StatusOK, isAuthorized: true},
		{name: "Reject request with an unauthorised user", request: newRequestWithAuthHeader(t, fmt.Sprintln("Bearer ", test_helpers.RandomString())), expectedStatusCode: http.StatusUnauthorized, isAuthorized: false},
		{name: "Reject request with an empty authorisation header", request: newRequest(t), expectedStatusCode: http.StatusUnauthorized, isAuthorized: false},
		{name: "Reject request without Bearer prefix", request: newRequestWithAuthHeader(t, testToken), expectedStatusCode: http.StatusUnauthorized, isAuthorized: false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var (
				stubCalled = false
				// stub handler is the next handler in the call chain
				stubHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					stubCalled = true
				})
				authMiddleware = NewAuthMiddleware(authorisedUsers)
				resRecorder    = httptest.NewRecorder()
			)
			authMiddleware.AuthHandler(stubHandler).ServeHTTP(resRecorder, tc.request)
			assert.Equal(t, tc.isAuthorized, stubCalled, "stub handler should never be called when request is unauthorized")
			assert.Equal(t, tc.expectedStatusCode, resRecorder.Code, "response status code does not match expected status code")
		})
	}
}

func newRequest(t *testing.T) *http.Request {
	request, err := http.NewRequest(http.MethodPost, "", http.NoBody)
	assert.NoError(t, err, "failed to create new http request")
	return request
}

func newRequestWithAuthHeader(t *testing.T, token string) *http.Request {
	req := newRequest(t)
	req.Header.Add("Authorization", token)
	return req
}
