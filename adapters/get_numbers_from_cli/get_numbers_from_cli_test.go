package get_numbers_from_cli

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetNumbersFromCLI(t *testing.T) {
	t.Skip()
	strOfNum := []string{"1,2,3"}
	want := []int{1, 2, 3}
	got := GetNumbersFromCLI(strOfNum)

	assert.Equal(t, want, got)
}
