package data_distributor

import (
	"fmt"
)

func DataDistributor(
	ArrayOfFileNamesFromCLI []string,
	ArrayOfNumbersFromCLI []string,
	extractAndDeduplicateNumbers func([]string) []int,
	getNumbersFromFile func([]string) []int,
) []int {
	var numbers []int
	fmt.Println(ArrayOfFileNamesFromCLI)
	if len(ArrayOfNumbersFromCLI) != 0 && len(ArrayOfFileNamesFromCLI) != 0 {
		return nil
	} else if len(ArrayOfNumbersFromCLI) != 0 {
		numbers = extractAndDeduplicateNumbers(ArrayOfNumbersFromCLI)
	} else if len(ArrayOfFileNamesFromCLI) != 0 {
		numbers = getNumbersFromFile(ArrayOfFileNamesFromCLI)
	}
	return numbers
}
