package printer

import (
	"bytes"
	"testing"
)

func TestPrinter(t *testing.T) {
	t.Run("Given a sum not equal to zero should return success message", func(t *testing.T) {
		expect := "Success, sum is 123\n"

		var output bytes.Buffer

		Printer(&output, "123")
		if expect != output.String() {
			t.Errorf("got %s expected %s", output.String(), expect)
		}
	})

	t.Run("Given a sum equal to zero should return unsuccessful message", func(t *testing.T) {
		expect := "Unsuccessful! Sum is 0\n. Please provide either file names or numbers, but not both.\n"

		var output bytes.Buffer

		Printer(&output, "0")
		if expect != output.String() {
			t.Errorf("got %s expected %s", output.String(), expect)
		}
	})
}
