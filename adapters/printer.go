package adapters

import (
	"fmt"
	"io"
)

func Printer(w io.Writer, formattedSum string) {
	if formattedSum == "0" {
		fmt.Fprintf(w, "Unsuccessful! Sum is %v\n. Please provide either file names or numbers, but not both.\n", formattedSum)
	} else {
		fmt.Fprintf(w, "Success, sum is %v\n", formattedSum)
	}
}
