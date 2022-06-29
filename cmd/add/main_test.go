package main_test

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
	t.Skip()
	file := "input2.csv"

	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	cmdPath := filepath.Join(dir, binName)
	fmt.Printf("here is it %s\n", cmdPath)

	cmd := exec.Command(cmdPath, "--input-file", file)

	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("ops.. seems like the flag name is wrong")
		t.Fatal(err)
	}

	expected := fmt.Sprintf("Success, sum is 867,685\n")

	if expected != string(out) {
		t.Errorf("Expected %q, got %q instead\n", expected, string(out))
	}
}
