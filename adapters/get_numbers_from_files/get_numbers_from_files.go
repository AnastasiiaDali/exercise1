package get_numbers_from_files

import (
	"fmt"
	"os"
	"strings"

	"exercise1/helpers/data_converter"
	"exercise1/helpers/unique_numbers"
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

	tempNumbers = data_converter.DataConverter(strings.Join(arrayOfInputs[:], ","))

	numbers := unique_numbers.UniqueNumbers(tempNumbers)

	return numbers
}
