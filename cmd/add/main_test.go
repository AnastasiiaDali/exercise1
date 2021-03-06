package main_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
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

func TestCLIFlags(t *testing.T) {
	numbers := "1,2,3"

	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	cmdPath := filepath.Join(dir, binName)

	cmd := exec.Command(cmdPath, "--input-numbers", numbers)

	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("ops.. seems like the flag name is wrong")
		t.Fatal(err)
	}

	expected := fmt.Sprintf("Success, sum is 6\n")

	if expected != string(out) {
		t.Errorf("Expected %q, got %q instead\n", expected, string(out))
	}

}

func TestCLIFlags2(t *testing.T) {
	file, err := ioutil.TempFile(".", "example.txt")
	require.NoError(t, err)

	_, err = file.Write([]byte("1,867683"))
	require.NoError(t, err)

	defer os.RemoveAll(file.Name())

	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	cmdPath := filepath.Join(dir, binName)
	fmt.Printf("here is it %s\n", cmdPath)

	cmd := exec.Command(cmdPath, "--input-file", file.Name())

	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("ops.. seems like the flag name is wrong")
		t.Fatal(err)
	}

	expected := fmt.Sprintf("Success, sum is 867,684\n")

	if expected != string(out) {
		t.Errorf("Expected %q, got %q instead\n", expected, string(out))
	}
}
