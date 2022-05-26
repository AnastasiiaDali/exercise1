package helpers

import (
	"regexp"
)

func DataConverter(input string) []int {
	var arrayOfnumbers []string
	var numbersFromFile []int

	a := regexp.MustCompile(`(\s*(,|\n)\s*)`)
	arrayOfnumbers = a.Split(input, -1)

	numbersFromFile = StringToIntConverter(arrayOfnumbers)

	return numbersFromFile
}
