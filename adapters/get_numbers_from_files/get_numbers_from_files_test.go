package get_numbers_from_files_test

import (
	"testing"

	"exercise1/adapters/get_numbers_from_files"
	"github.com/stretchr/testify/assert"
)

//get number from the file function returns an array of integers extracted from the file.
func TestGetNumbersFromFile(t *testing.T) {
	t.Run("should return array of integers from one file", func(t *testing.T) {
		testArrayOfFileName1 := []string{"input1.csv"}

		expected := []int{1, 2, 32, 321, 1234, 4567890}
		result := get_numbers_from_files.GetNumbersFromFile(testArrayOfFileName1)
		assert.Equal(t, expected, result)
	})

	//numbers in test_numbers2.txt and test_numbers.csv file are the same,
	//number from the file function should remove duplicated numbers

	t.Run("should return array of integers from three file", func(t *testing.T) {
		t.Skip()
		testArrayOfFileName2 := []string{"/Users/anastasiia.dalakishvili/github/personal/exercise1/testdata/test_numbers.csv", "/Users/anastasiia.dalakishvili/github/personal/exercise1/testdata/test_numbers.txt", "/Users/anastasiia.dalakishvili/github/personal/exercise1/testdata/test_number2.txt"}

		expected := []int{1, 2, 32, 321, 1234, 4567890, 33, 44, 55, 678}
		result := get_numbers_from_files.GetNumbersFromFile(testArrayOfFileName2)
		assert.Equal(t, expected, result)
	})
}
