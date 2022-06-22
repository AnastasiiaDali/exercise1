package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"os/exec"
	"testing"

	http2 "exercise1/adapters/http"
)

var (
	binName = "math"
)

func TestMain(m *testing.M) {
	fmt.Println("Building tool...")
	build := exec.Command("go", "build", "-o", binName)

	if err := build.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Cannot build tool %s: %s", binName, err)
		os.Exit(1)
	}

	fmt.Println("Running tests....")
	result := m.Run()

	fmt.Println("Cleaning up...")
	os.Remove(binName)

	os.Exit(result)
}

func TestServer(t *testing.T) {
	t.Run("returns the sum", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "add?num=76&num=65", nil)
		response := httptest.NewRecorder()

		http2.MathHandler(response, request)
		//http.MathHandler(response, request)

		got := response.Body.String()
		want := "141"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
