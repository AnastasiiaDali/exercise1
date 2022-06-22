package data_distributor

import (
	"fmt"

	"exercise1/adapters/get_numbers_from_cli"
	"exercise1/adapters/get_numbers_from_files"
)

func DataDistributor(ArrayOfFileNamesFromCLI []string, ArrayOfNumbersFromCLI []string) []int {
	var numbers []int
	fmt.Println(ArrayOfFileNamesFromCLI)
	if len(ArrayOfNumbersFromCLI) != 0 && len(ArrayOfFileNamesFromCLI) != 0 {
		return nil
	} else if len(ArrayOfNumbersFromCLI) != 0 {
		numbers = get_numbers_from_cli.GetNumbersFromCLI(ArrayOfNumbersFromCLI)
	} else if len(ArrayOfFileNamesFromCLI) != 0 {
		numbers = get_numbers_from_files.GetNumbersFromFile(ArrayOfFileNamesFromCLI)
	}
	return numbers
}
