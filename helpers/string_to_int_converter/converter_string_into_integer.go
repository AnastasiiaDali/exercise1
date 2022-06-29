package string_to_int_converter

import "strconv"

func StringToIntConverter(arrayOfStrings []string) []int {
	var arrayOfInt []int
	for _, number := range arrayOfStrings {
		var integer int
		if number == "" {
			integer = 0
		} else {
			integer, _ = strconv.Atoi(number)
		}
		arrayOfInt = append(arrayOfInt, integer)
	}
	return arrayOfInt
}
