package data_distributor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func mockExtractAndDeduplicateNumbers([]string) []int {
	return nil
}
func mockGetNumbersFromFile([]string) []int {
	return nil
}

func TestDataDistributor(t *testing.T) {
	t.Skip()
	t.Run("Given array of file names should call GetNumbersFromFile function", func(t *testing.T) {
		fakeArrayOfFileNames := []string{"file1", "file2"}
		fakeArrayOfNumbersFromCli := []string{}

		DataDistributor(fakeArrayOfFileNames, fakeArrayOfNumbersFromCli, mockExtractAndDeduplicateNumbers, mockGetNumbersFromFile)

		assert.Len(t, mockExtractAndDeduplicateNumbers, 1)
		assert.Len(t, mockGetNumbersFromFile, 0)
	})
}

//mocks
//create a mock for func and pass mock func to datadist and assert
