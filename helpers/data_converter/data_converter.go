package data_converter

import (
	"regexp"

	"exercise1/helpers/string_to_int_converter"
)

func DataConverter(input string) []int {
	var arrayOfnumbers []string
	var numbersFromFile []int

	a := regexp.MustCompile(`(\s*(,|\n)\s*)`)
	arrayOfnumbers = a.Split(input, -1)

	numbersFromFile = string_to_int_converter.StringToIntConverter(arrayOfnumbers)

	return numbersFromFile
}
