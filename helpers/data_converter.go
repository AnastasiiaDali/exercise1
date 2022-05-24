package helpers

import (
	"regexp"
	"strconv"
)

func DataConverter(input string) []int {
	var arrayOfnumbers []string
	var numbersFromFile []int

	a := regexp.MustCompile(`(\s*(,|\n)\s*)`)
	arrayOfnumbers = a.Split(input, -1)

	for _, number := range arrayOfnumbers {
		integer, _ := strconv.Atoi(number)
		numbersFromFile = append(numbersFromFile, integer)
	}

	return numbersFromFile
}
