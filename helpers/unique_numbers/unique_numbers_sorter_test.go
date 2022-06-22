package unique_numbers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniqueNumbers(t *testing.T) {
	t.Skip()
	unsortedNumber := []int{1, 2, 3, 4, 5, 5}
	want := []int{1, 2, 3, 4, 5}
	got := UniqueNumbers(unsortedNumber)

	assert.Equal(t, want, got)
}
