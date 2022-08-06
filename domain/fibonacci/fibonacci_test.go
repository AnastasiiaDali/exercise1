package fibonacci

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_FibonacciNumber(t *testing.T) {
	testCases := []struct {
		input  int
		output string
	}{
		{input: 0, output: "0"},
		{input: 9, output: "34"},
		{input: 15, output: "610"},
	}

	for _, tc := range testCases {
		fibonacci := NewFibonacci()
		result, err := fibonacci.FibonacciNumber(tc.input)
		assert.NoError(t, err)
		assert.Equal(t, tc.output, result)
	}

	testCasesWithErrors := []struct {
		input  int
		output string
	}{
		{input: 93, output: "0"},
		{input: -3, output: "0"},
	}

	for _, tc := range testCasesWithErrors {
		fibonacci := NewFibonacci()
		result, err := fibonacci.FibonacciNumber(tc.input)
		assert.Error(t, err)
		assert.Equal(t, tc.output, result)
	}
}
