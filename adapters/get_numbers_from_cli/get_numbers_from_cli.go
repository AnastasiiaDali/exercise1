package get_numbers_from_cli

import (
	"strconv"
	"strings"

	"exercise1/helpers"
)

func GetNumbersFromCLI(ArrayOfNumbersFromCLI []string) []int {
	var tempNumbers []int
	var arrayOfNumbers []string

	for _, str := range ArrayOfNumbersFromCLI {
		arrayOfNumbers = strings.Split(str, ",")
	}

	for _, i := range arrayOfNumbers {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		tempNumbers = append(tempNumbers, j)
	}
	numbers := helpers.UniqueNumbers(tempNumbers)
	return numbers
}
