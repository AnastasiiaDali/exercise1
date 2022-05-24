package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDataConverter(t *testing.T) {
	t.Run("Converter should return an array of integers", func(t *testing.T) {
		data := "1,2,3"
		var desiredRes []int
		result := DataConverter(data)

		assert.IsType(t, desiredRes, result)
	})
}
