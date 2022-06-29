package string_to_int_converter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringToIntConverter(t *testing.T) {
	testCases := []struct {
		input []string
		want  []int
	}{
		{input: []string{"1", "2", "3"}, want: []int{1, 2, 3}},
		{input: []string{"11", "", ""}, want: []int{11, 0, 0}},
		{input: []string{"11", "hi", "there"}, want: []int{11, 0, 0}},
	}

	for _, tc := range testCases {
		t.Run("Converter should return an array of integers when given array of strings", func(t *testing.T) {
			result := StringToIntConverter(tc.input)
			assert.Equal(t, tc.want, result)
		})
	}
}
