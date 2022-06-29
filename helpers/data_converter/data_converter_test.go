package data_converter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDataConverter(t *testing.T) {
	testCases := []struct {
		input string
		want  []int
	}{
		{input: "1,2,3", want: []int{1, 2, 3}},
		{input: "1\n2\n3", want: []int{1, 2, 3}},
	}

	for _, tc := range testCases {
		t.Run("Converter should return an array of integers", func(t *testing.T) {
			result := DataConverter(tc.input)

			assert.IsType(t, result, tc.want)
		})
	}
}
