package printer

import (
	"bytes"
	"testing"
)

func TestPrinter(t *testing.T) {

	t.Run("testing", func(t *testing.T) {
		expect := "Success, sum is 123\n"

		var output bytes.Buffer

		Printer(&output, "123")
		if expect != output.String() {
			t.Errorf("got %s expected %s", output.String(), expect)
		}
	})

	t.Run("testing", func(t *testing.T) {
		expect := "Unsuccessful! Sum is 0\n. Please provide either file names or numbers, but not both.\n"

		var output bytes.Buffer

		Printer(&output, "0")
		if expect != output.String() {
			t.Errorf("got %s expected %s", output.String(), expect)
		}
	})
}
