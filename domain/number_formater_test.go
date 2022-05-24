package domain

import (
	"testing"
)

func TestFormatNumber(t *testing.T) {
	t.Run("Should return number without commas", func(t *testing.T) {
		number := 9999

		got := FormatNumber(number)
		want := "9999"

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("Should return number with commas", func(t *testing.T) {
		number := 10000

		got := FormatNumber(number)
		want := "10,000"

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
