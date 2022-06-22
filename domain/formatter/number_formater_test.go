package formatter

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

	t.Run("Given large number should return correctly formatted number with commas", func(t *testing.T) {
		number := 100000000000

		got := FormatNumber(number)
		want := "100,000,000,000"

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
