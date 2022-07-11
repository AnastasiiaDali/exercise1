package data_distributor_test

import (
	"io/fs"
	"testing"

	"exercise1/adapters/data_distributor"
	"github.com/stretchr/testify/assert"
)

type mockDataCollector struct {
}

func (dc *mockDataCollector) ExtractAndDeduplicateNumbersFromFiles2(fs fs.FS) []int {
	//TODO implement me
	return []int{}
}

func (dc *mockDataCollector) ExtractAndDeduplicateNumbersFromCLI([]string) []int {
	var numbers []int
	return numbers
}

func (dc *mockDataCollector) ExtractAndDeduplicateNumbersFromFiles([]string) []int {
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
