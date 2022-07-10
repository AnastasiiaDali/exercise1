package calculator_test

import (
	"math"
	"testing"

	"exercise1/domain/calculator"
)

func TestAdd(t *testing.T) {
	testCases := []struct {
		input []int
		want  int
	}{
		{input: []int{5}, want: 5},
		{input: []int{5, 4, 2, -10, 32, 14}, want: 47},
		{input: []int{}, want: 0},
		{input: []int{math.MaxInt, 2}, want: 0},
	}
	for _, tc := range testCases {
		t.Run("Should take any number of integers and print out the sum", func(t *testing.T) {
			calculator := calculator.New()
			got := calculator.Add(tc.input)

			if got != tc.want {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}
