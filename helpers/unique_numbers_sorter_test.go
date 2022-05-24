package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniqueNumbers(t *testing.T) {
	unsortedNumber := []int{1, 2, 3, 4, 5, 5}
	want := []int{1, 2, 3, 4, 5}
	got := UniqueNumbers(unsortedNumber)

	assert.Equal(t, want, got)
}
