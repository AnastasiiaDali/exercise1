package get_numbers_from_files

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

//get number from the file function returns an array of integers extracted from the file.
func TestGetNumbersFromFile(t *testing.T) {
	t.Run("should return a string of numbers separated by comma - imitate txt file", func(t *testing.T) {
		expectedData := "1, 2, 3, 4"

		var buffer bytes.Buffer

		buffer.WriteString("1, 2, 3, 4")
		content, err := readFile(&buffer)
		if err != nil {
			t.Error("Failed to read csv data")
		}
		assert.Equal(t, content, expectedData)

	})
	t.Run("should return a string of numbers, each number on the new line - imitate csv file", func(t *testing.T) {
		expectedData := "1\n2\n3\n4\n500"

		var buffer bytes.Buffer

		buffer.WriteString("1\n2\n3\n4\n500")
		content, err := readFile(&buffer)
		if err != nil {
			t.Error("Failed to read csv data")
		}
		assert.Equal(t, content, expectedData)
	})
}
