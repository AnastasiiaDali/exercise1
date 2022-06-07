package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

var (
	binName = "add"
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

//go run cmd/add/main.go --input-numbers="1,2,3"

func TestCLI(t *testing.T) {
	numbers := "1,2,3"
	file := "data/input2.csv"

	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	cmdPath := filepath.Join(dir, binName)

	t.Run("Add number from CLI", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "--input-numbers", numbers)

		if err := cmd.Run(); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Add file names from CLI", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "--input-file", file)

		if err := cmd.Run(); err != nil {
			t.Fatal(err)
		}
	})
}
