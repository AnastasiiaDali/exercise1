package formatter_test

import (
	"testing"

	formatter "exercise1/domain/formatter"
)

func TestFormatNumber(t *testing.T) {
	formatter := formatter.New()

	t.Run("Given the number smaller than 10000 should return number without commas", func(t *testing.T) {
		number := 9999

		got := formatter.FormatNumbers(number)
		want := "9999"

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("Given the number greater or equal to 10000 should return number with commas", func(t *testing.T) {
		number := 10000

		got := formatter.FormatNumbers(number)
		want := "10,000"

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("Given large number should return correctly formatted number with commas", func(t *testing.T) {
		number := 100000000000

		got := formatter.FormatNumbers(number)
		want := "100,000,000,000"

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
