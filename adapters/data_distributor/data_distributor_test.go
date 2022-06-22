package data_distributor_test

import (
	"testing"

	"exercise1/adapters/data_distributor"
	"github.com/stretchr/testify/assert"
)

func TestDataDistributor(t *testing.T) {
	t.Skip()
	t.Run("test with numbers passed to the distributor", func(t *testing.T) {
		mockArrayOfFileNamesFromCLI := []string{}
		mockArrayOfNumbersFromCLI := []string{"1,2,3,4"}
		result := data_distributor.DataDistributor(mockArrayOfFileNamesFromCLI, mockArrayOfNumbersFromCLI)
		expected := []int{1, 2, 3, 4}
		assert.Equal(t, expected, result)
	})
	t.Run("test with file name passed to the distributor", func(t *testing.T) {
		mockArrayOfFileNamesFromCLI := []string{"input1.txt"}
		mockArrayOfNumbersFromCLI := []string{}
		result := data_distributor.DataDistributor(mockArrayOfFileNamesFromCLI, mockArrayOfNumbersFromCLI)
		expected := []int{1, 2, 32, 321, 1234, 4567890}
		assert.Equal(t, expected, result)
	})

	t.Run("test with 2 file names passed to the distributor", func(t *testing.T) {
		mockArrayOfFileNamesFromCLI := []string{"input2.csv", "input1.txt"}
		mockArrayOfNumbersFromCLI := []string{}
		result := data_distributor.DataDistributor(mockArrayOfFileNamesFromCLI, mockArrayOfNumbersFromCLI)
		expected := []int{1, 2, 32, 321, 1234, 4567890, 33, 44, 55, 678}
		assert.Equal(t, expected, result)
	})
}
