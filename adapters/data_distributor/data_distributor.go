package data_distributor

func DataDistributor(
	ArrayOfFileNamesFromCLI []string,
	ArrayOfNumbersFromCLI []string,
	extractAndDeduplicateNumbers func([]string) []int,
	getNumbersFromFile func([]string) []int,
) []int {
	var numbers []int

	if len(ArrayOfNumbersFromCLI) != 0 && len(ArrayOfFileNamesFromCLI) != 0 {
		return nil
	} else if len(ArrayOfNumbersFromCLI) != 0 {
		numbers = extractAndDeduplicateNumbers(ArrayOfNumbersFromCLI)
	} else if len(ArrayOfFileNamesFromCLI) != 0 {
		numbers = getNumbersFromFile(ArrayOfFileNamesFromCLI)
	}

	return numbers
}
