package get_numbers_from_cli_test

import (
	"testing"

	"exercise1/adapters/get_numbers_from_cli"
	"github.com/stretchr/testify/assert"
)

func TestExtractAndDeduplicateNumbers(t *testing.T) {
	testCases := []struct {
		input []string
		want  []int
	}{
		{input: []string{"1,2,3"}, want: []int{1, 2, 3}},
		{input: []string{"1,2,2,3,3"}, want: []int{1, 2, 3}},
		{input: []string{"1,2,3,l"}, want: []int{1, 2, 3, 0}},
		{input: []string{"3.3,1,2,3,l,4,l"}, want: []int{0, 1, 2, 3, 4}},
		{input: []string{"1, 2, 3 "}, want: []int{1, 2, 3}},
	}
	for _, tc := range testCases {
		result := get_numbers_from_cli.ExtractAndDeduplicateNumbers(tc.input)
		assert.Equal(t, tc.want, result)
	}
}
