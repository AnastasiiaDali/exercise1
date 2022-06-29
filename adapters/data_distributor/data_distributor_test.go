package data_distributor

import (
	"testing"

	"exercise1/adapters/get_numbers_from_cli"
	"exercise1/adapters/get_numbers_from_files"
	"github.com/stretchr/testify/assert"
)

func TestDataDistributor(t *testing.T) {
	t.Skip()
	t.Run("Given array of file names should call GetNumbersFromFile function", func(t *testing.T) {
		fakeArrayOfFileNames := []string{"file1", "file2"}
		fakeArrayOfNumbersFromCli := []string{}

		DataDistributor(fakeArrayOfFileNames, fakeArrayOfNumbersFromCli, get_numbers_from_cli.ExtractAndDeduplicateNumbers, get_numbers_from_files.GetNumbersFromFile)

		assert.Len(t, get_numbers_from_files.GetNumbersFromFile(fakeArrayOfFileNames), 1)
		assert.Len(t, get_numbers_from_cli.ExtractAndDeduplicateNumbers(fakeArrayOfNumbersFromCli), 0)
	})
}
