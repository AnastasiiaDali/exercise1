package get_numbers_from_cli_test

import (
	"testing"

	"exercise1/adapters/get_numbers_from_cli"
	"github.com/stretchr/testify/assert"
)

func TestExtractAndDeduplicateNumbers(t *testing.T) {
	strOfNum := []string{"1,2,3"}
	want := []int{1, 2, 3}
	got := get_numbers_from_cli.ExtractAndDeduplicateNumbers(strOfNum)

	assert.Equal(t, want, got)
}
