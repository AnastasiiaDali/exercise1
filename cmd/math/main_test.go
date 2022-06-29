package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	http2 "exercise1/adapters/http"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
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

func TestCLIArg(t *testing.T) {
	t.Skip()
	fmt.Println("I am here starting test")
	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	cmdPath := filepath.Join(dir, binName)

	t.Run("start the server after passing correct flag", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "--web-server")

		fmt.Println("I am about to run")
		err := cmd.Run()
		if err != nil {
			fmt.Println("error executing the program")
			t.Fatal(err)
		}
		fmt.Println("I am running now")

		r := mux.NewRouter()
		r.HandleFunc("/", http2.HealthCheckHandler).Methods(http.MethodGet)

		req := httptest.NewRequest(http.MethodGet, "/wedfew", nil)
		res := httptest.NewRecorder()

		http2.HealthCheckHandler(res, req)

		assert.Equal(t, http.StatusOK, res.Code)
		cmd.Process.Kill()
	})
}
