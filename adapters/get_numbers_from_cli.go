package adapters

import (
	"strconv"
	"strings"
)

func GetNumbersFromCLI() []int {
	var numbers []int
	var arrayOfNumbers []string

	for _, str := range ArrayOfNumbersFromCLI {
		arrayOfNumbers = strings.Split(str, ",")
	}

	for _, i := range arrayOfNumbers {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, j)
	}

	return numbers
}
