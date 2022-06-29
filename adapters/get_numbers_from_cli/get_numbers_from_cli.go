package get_numbers_from_cli

import (
	"fmt"
	"strconv"
	"strings"

	"exercise1/helpers/unique_numbers"
)

func ExtractAndDeduplicateNumbers(ArrayOfNumbersFromCLI []string) []int {
	var tempNumbers []int
	var arrayOfNumbers []string

	for _, str := range ArrayOfNumbersFromCLI {
		stringWithoutSpaces := strings.ReplaceAll(str, " ", "")
		arrayOfNumbers = strings.Split(stringWithoutSpaces, ",")
	}

	for _, i := range arrayOfNumbers {
		j, err := strconv.Atoi(i)
		if err != nil {
			j = 0
			fmt.Printf("not able to convert the %v that is not an integer. zero value has been assign instead\n", i)
		}
		tempNumbers = append(tempNumbers, j)
	}
	numbers := unique_numbers.UniqueNumbers(tempNumbers)
	return numbers
}
