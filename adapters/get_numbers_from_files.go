package adapters

import (
	"fmt"
	"os"
	"strings"

	"exercise1/helpers"
)

func GetNumbersFromFile(ArrayOfFileNamesFromCLI []string) []int {
	var tempNumbers []int
	var arrayOfInputs []string

	for _, file := range ArrayOfFileNamesFromCLI {
		data, err := os.ReadFile(file)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		input := string(data)
		arrayOfInputs = append(arrayOfInputs, input)
	}

	tempNumbers = helpers.DataConverter(strings.Join(arrayOfInputs[:], ","))

	numbers := helpers.UniqueNumbers(tempNumbers)

	return numbers
}
