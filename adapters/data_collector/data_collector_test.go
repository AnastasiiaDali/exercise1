package data_collector_test

import (
	"bytes"
	"testing"

	"exercise1/adapters/data_collector"
	"github.com/stretchr/testify/assert"
)

func TestExtractAndDeduplicateNumbersFromCLI(t *testing.T) {
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

	dataCollector := data_collector.New()
	for _, tc := range testCases {
		result := dataCollector.ExtractAndDeduplicateNumbersFromCLI(tc.input)
		assert.Equal(t, tc.want, result)
	}
}

func TestExtractAndDeduplicateNumbersFromFiles(t *testing.T) {
	t.Run("should return a string of numbers separated by comma - imitate txt file", func(t *testing.T) {
		expectedData := "1, 2, 3, 4"

		var buffer bytes.Buffer

		buffer.WriteString("1, 2, 3, 4")
		content, err := data_collector.ReadFile(&buffer)
		if err != nil {
			t.Error("Failed to read csv data")
		}
		assert.Equal(t, content, expectedData)

	})
	t.Run("should return a string of numbers, each number on the new line - imitate csv file", func(t *testing.T) {
		expectedData := "1\n2\n3\n4\n500"

		var buffer bytes.Buffer

		buffer.WriteString("1\n2\n3\n4\n500")
		content, err := data_collector.ReadFile(&buffer)
		if err != nil {
			t.Error("Failed to read csv data")
		}
		assert.Equal(t, content, expectedData)
	})
}
