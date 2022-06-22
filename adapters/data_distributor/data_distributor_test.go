package data_distributor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDataDistributor(t *testing.T) {

	t.Run("test with numbers passed to the distributor", func(t *testing.T) {
		mockArrayOfFileNamesFromCLI := []string{}
		mockArrayOfNumbersFromCLI := []string{"1,2,3,4"}
		result := DataDistributor(mockArrayOfFileNamesFromCLI, mockArrayOfNumbersFromCLI)
		expected := []int{1, 2, 3, 4}
		assert.Equal(t, expected, result)
	})
	t.Run("test with file name passed to the distributor", func(t *testing.T) {
		mockArrayOfFileNamesFromCLI := []string{"/Users/anastasiia.dalakishvili/github/personal/exercise1/testdata/test_numbers.csv"}
		mockArrayOfNumbersFromCLI := []string{}
		result := DataDistributor(mockArrayOfFileNamesFromCLI, mockArrayOfNumbersFromCLI)
		expected := []int{1, 2, 32, 321, 1234, 4567890}
		assert.Equal(t, expected, result)
	})

	t.Run("test with 2 file names passed to the distributor", func(t *testing.T) {
		mockArrayOfFileNamesFromCLI := []string{"/Users/anastasiia.dalakishvili/github/personal/exercise1/testdata/test_numbers.csv", "/Users/anastasiia.dalakishvili/github/personal/exercise1/testdata/test_numbers.txt"}
		mockArrayOfNumbersFromCLI := []string{}
		result := DataDistributor(mockArrayOfFileNamesFromCLI, mockArrayOfNumbersFromCLI)
		expected := []int{1, 2, 32, 321, 1234, 4567890, 33, 44, 55, 678}
		assert.Equal(t, expected, result)
	})
}
