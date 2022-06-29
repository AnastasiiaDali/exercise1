package unique_numbers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniqueNumbers(t *testing.T) {
	testCases := []struct {
		input []int
		want  []int
	}{
		{input: []int{1, 2, 3, 4, 5, 5}, want: []int{1, 2, 3, 4, 5}},
		{input: []int{1, 1, 1}, want: []int{1}},
	}
	for _, tc := range testCases {
		t.Run("given an array of number should only unique numbers", func(t *testing.T) {
			result := UniqueNumbers(tc.input)
			assert.Equal(t, tc.want, result)
		})
	}
}
