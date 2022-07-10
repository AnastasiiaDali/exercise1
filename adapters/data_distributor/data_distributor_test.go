package data_distributor_test

import (
	"testing"

	"exercise1/adapters/data_distributor"
	"github.com/stretchr/testify/assert"
)

type mockDataCollector struct {
}

func (dc *mockDataCollector) ExtractAndDeduplicateNumbersFromCLI(numbersFromCLI []string) []int {
	var numbers []int
	return numbers
}

func (dc *mockDataCollector) ExtractAndDeduplicateNumbersFromFiles(fileNamesFromCLI []string) []int {
	var numbers []int
	return numbers
}

func TestDistribute(t *testing.T) {
	t.Skip()
	t.Run("Given array of file names should call ExtractAndDeduplicateNumbersFromFiles function", func(t *testing.T) {
		fakeFileNames := []string{"file1", "file2"}
		fakeNumbersFromCli := []string{}

		dataCollector := &mockDataCollector{}
		dataDistributor := data_distributor.New(dataCollector)

		dataDistributor.Distribute(fakeFileNames, fakeNumbersFromCli)

		assert.Len(t, dataCollector.ExtractAndDeduplicateNumbersFromFiles, 1)
		assert.Len(t, dataCollector.ExtractAndDeduplicateNumbersFromCLI, 0)
	})
}

//
//func mockExtractAndDeduplicateNumbers([]string) []int {
//	return nil
//}
//func mockGetNumbersFromFile([]string) []int {
//	return nil
//}
//
//func TestDataDistributor(t *testing.T) {
//	t.Skip()
//	t.Run("Given array of file names should call ExtractAndDeduplicateNumbersFromFiles function", func(t *testing.T) {
//		fakeArrayOfFileNames := []string{"file1", "file2"}
//		fakeArrayOfNumbersFromCli := []string{}
//
//		DataDistributor(fakeArrayOfFileNames, fakeArrayOfNumbersFromCli, mockExtractAndDeduplicateNumbers, mockGetNumbersFromFile)
//
//		assert.Len(t, mockExtractAndDeduplicateNumbers, 1)
//		assert.Len(t, mockGetNumbersFromFile, 0)
//	})
//}

//mocks
//create a mock for func and pass mock func to datadist and assert
