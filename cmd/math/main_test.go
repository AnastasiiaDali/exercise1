package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
	"time"

	temphttp "exercise1/adapters/http"
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

func TestMath_CalculatesAndReturnsSum(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	cmdPath := filepath.Join(dir, binName)
	cmd := exec.Command(cmdPath)
	assert.NoError(t, cmd.Start())

	waitForServer()

	client := temphttp.NewClient("http://localhost:8081", &http.Client{})

	numbers := []string{"11", "10"}
	temp := client.Convert(numbers)

	assert.NoError(t, err)
	assert.Equal(t, "21", temp)
}

func waitForServer() {
	for i := 0; i < 10; i++ {
		conn, _ := net.Dial("tcp", net.JoinHostPort("localhost", "8081"))
		if conn != nil {
			conn.Close()
			break
		}
		time.Sleep(100 * time.Millisecond)
	}
}
