package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
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

func TestCLI(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	cmdPath := filepath.Join(dir, binName)

	t.Run("start the server after passing correct flag", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "--web-server")

		if err := cmd.Run(); err != nil {
			t.Fatal(err)
		}
	})
}
