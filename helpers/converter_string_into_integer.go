package helpers

import "strconv"

func StringToIntConverter(arrayOfStrings []string) []int {
	var arrayOfInt []int
	for _, number := range arrayOfStrings {
		integer, _ := strconv.Atoi(number)
		arrayOfInt = append(arrayOfInt, integer)
	}
	return arrayOfInt
}
