package string_to_int_converter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringToIntConverter(t *testing.T) {
	t.Run("Converter should return an array of integers when given array of strings", func(t *testing.T) {
		data := []string{"1", "2", "3"}
		var desiredRes []int
		result := StringToIntConverter(data)
		assert.IsType(t, desiredRes, result)
	})
}
